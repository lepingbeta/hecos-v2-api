package CodeLogin

import (
	"fmt"
	"os"

	dhhttp "github.com/lepingbeta/go-common-v2-dh-http"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	middleware "github.com/lepingbeta/go-common-v2-dh-middleware"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tangxiaoer.shop/dahe/hecos-v2-api/config"
	// {{占位符 import}}
)

func (p *CodeLogin) CodeLogin() {

	p.Code2Login()
	if p.Err != nil {
		return
	}
	p.AddJwtToken()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *CodeLogin) Code2Login() {

	// 定义要POST的URL
	urlPrefix := os.Getenv("OAUTH2_URL_PREFIX")
	urlCode2token := urlPrefix + "/oauth2/code2_token"
	urlAt2userinfo := urlPrefix + "/oauth2/oauth2_get_userinfo"
	urlRefreshToken := urlPrefix + "/oauth2/refresh_token"
	dhlog.DebugAny(p.Filter)
	// 定义要发送的数据
	reqTokenParams := map[string]interface{}{
		"code": p.Filter["code"],
	}

	var respToken map[string]any
	respToken, p.Err = dhhttp.PostJSON2Map(urlCode2token, reqTokenParams)
	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg(p.Err.Error())
		p.MsgKey = "user_code_login_Code2Login_code2_token_failed"
		return
		// return filter, resultMap, err.Error(), respToken["msg_key"].(string), err
	}

	if respToken["status"].(string) != "success" {
		dhlog.Error(respToken["msg"].(string))
		p.Err = fmt.Errorf(respToken["msg"].(string))
		p.Msg = utils.DebugMsg(p.Err.Error())
		p.MsgKey = "user_code_login_Code2Login_code2_token_error_status"
		return
	}

	respTokenData := respToken["data"].(map[string]any)
	accessToken := respTokenData["access_token"].(string)
	refreshToken := respTokenData["refresh_token"].(string)

	// 定义要发送的数据
	reqUserinfoParams := map[string]interface{}{
		"access_token": accessToken,
	}
	var respUserinfo map[string]any
	respUserinfo, p.Err = dhhttp.PutJSON2Map(urlAt2userinfo, reqUserinfoParams)
	dhlog.DebugAny(respUserinfo)
	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		// return filter, resultMap, p.Err.Error(), respUserinfo["msg_key"].(string), err
		p.Msg = utils.DebugMsg(p.Err.Error())
		p.MsgKey = "user_code_login_Code2Login_oauth2_get_userinfo_error"
		return
	}

	if respUserinfo["status"].(string) != "success" {
		// access token 失效，尝试用refresh token刷新
		if respUserinfo["msg_key"].(string) == "oauth2_oauth2_get_userinfo_getUserIdByAt_get_userid_failed" {
			reqNewTokenParams := map[string]interface{}{
				"refresh_token": refreshToken,
			}
			var respNewToken map[string]any
			respNewToken, p.Err = dhhttp.PostJSON2Map(urlRefreshToken, reqNewTokenParams)
			if p.Err != nil {
				dhlog.Error(p.Err.Error())
				// return filter, resultMap, p.Err.Error(), respUserinfo["msg_key"].(string), err
				p.Msg = utils.DebugMsg(p.Err.Error())
				p.MsgKey = "user_code_login_Code2Login_refresh_token_error"
				return
			}
			respNewTokenData := respNewToken["data"].(map[string]any)
			accessToken = respNewTokenData["access_token"].(string)
			refreshToken = respNewTokenData["refresh_token"].(string)

			// 定义要发送的数据
			reqUserinfoParams := map[string]interface{}{
				"access_token": accessToken,
			}
			respUserinfo, p.Err = dhhttp.PutJSON2Map(urlAt2userinfo, reqUserinfoParams)
			if p.Err != nil {
				dhlog.Error(p.Err.Error())
				// return filter, resultMap, p.Err.Error(), respUserinfo["msg_key"].(string), err
				p.Msg = utils.DebugMsg(p.Err.Error())
				p.MsgKey = "user_code_login_Code2Login_oauth2_get_userinfo_error2"
				return
			}
		} else {
			dhlog.Error(respUserinfo["msg"].(string))
			// return filter, resultMap, p.Err.Error(), respUserinfo["msg_key"].(string), err
			p.Msg = utils.DebugMsg(respUserinfo["msg"].(string))
			p.MsgKey = "user_code_login_Code2Login_oauth2_get_userinfo_error3"
			return
		}
	}

	// 查询用户是否已经存在
	respUserinfoData := respUserinfo["data"].(map[string]any)
	userFilter := bson.M{}
	userFilter["iuc_id"] = respUserinfoData["_id"].(string)
	var count int64
	count, p.Err = mongodb.Count("user", userFilter)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg(p.Err.Error())
		p.MsgKey = "user_code_login_Code2Login_query_user_error"
		return
	}

	// 查询不到此用户，创建一个新用户
	if count == 0 {
		userinfo := respUserinfoData
		userinfo["iuc_access_token"] = accessToken
		userinfo["iuc_refresh_token"] = refreshToken
		userinfo["iuc_id"] = respUserinfoData["_id"].(string)
		delete(userinfo, "_id")
		delete(userinfo, "create_time")
		delete(userinfo, "update_time")
		var userDbData bson.D
		userDbData, p.Err = mongodb.MapToBsonD(userinfo)
		if p.Err != nil {
			dhlog.Error(p.Err.Error())
			p.Msg = utils.DebugMsg("数据转成bson.D失败: " + p.Err.Error())
			p.MsgKey = "user_code_login_Code2Login_mapToBsonD_failed"
			return
		}

		_, p.Err = mongodb.InsertOneBsonD("user", userDbData)

		if p.Err != nil {
			dhlog.Error(p.Err.Error())
			p.Msg = utils.DebugMsg("数据转成bson.D失败: " + p.Err.Error())
			p.MsgKey = "user_code_login_Code2Login_InsertOneBsonD_failed"
			return

			// dhlog.Error(err.Error())
			// return userFilter, nil, "数据入库失败: " + err.Error(), "user_code_login_code2login_insert_to_db_failed", err
		}
	}

	fieldList := bson.D{
		{Key: "_id", Value: 1},
		{Key: "account", Value: 1},
	}
	// 创建Find选项，设置Projection
	findOptions := options.FindOne()
	findOptions.SetProjection(fieldList)
	p.Result, p.Err = mongodb.FindOne("user", userFilter, findOptions)
	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("FindOne 失败: " + p.Err.Error())
		p.MsgKey = "user_code_login_Code2Login_FindOne_failed"
		return
	}
}

func (p *CodeLogin) AddJwtToken() {

	p.Result.(bson.M)["token"], p.Result.(bson.M)["refreshToken"], p.Err = middleware.JWTGenerateToken(config.JwtSecret, config.JwtRefreshSecret, p.Result.(bson.M)["_id"].(primitive.ObjectID).Hex(), p.Result.(bson.M)["account"].(string), config.JwtExpSec, config.JwtRefreshSec)
	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg(p.Err.Error())
		p.MsgKey = "user_code_login_AddJwtToken_JWTGenerateToken_error"
		return
	}
}

// {{占位符 composition}}
