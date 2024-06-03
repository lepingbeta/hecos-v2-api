package CodeLogin

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	dhhttp "github.com/lepingbeta/go-common-v2-dh-http"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	middleware "github.com/lepingbeta/go-common-v2-dh-middleware"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tangxiaoer.shop/dahe/hecos-v2-api/config"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

func preProcessing(params t.CodeLoginParams, filterEmpty bson.M, resultEmpty any, c *gin.Context) (bson.M, any, string, string, error) {
	filter, _ := utils.Struct2BsonM(params)
	var result = bson.M{}
	var msg, msgKey string = "", ""
	var err error = nil

	filter, result, msg, msgKey, err = code2login(params, filter, result, c)
	if err != nil {
		return filter, result, msg, msgKey, err
	}
	// {{占位符 preProcessing}}

	return filter, result, msg, msgKey, err
}

func CodeLogin(params t.CodeLoginParams, c *gin.Context) (any, string, string, error) {
	filter, finalResult, msg, msgKey, err := preProcessing(params, nil, nil, c)
	// {{bsonD holder}}
	if err != nil {
		return finalResult, msg, msgKey, err
	}

	filter, msg, msgKey, err = CodeLoginPre(filter, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	// fieldList := bson.D{{{fieldInList}}}
	// 创建Find选项，设置Projection
	findOptions := options.FindOne()
	// findOptions.SetProjection(fieldList)
	result, err := mongodb.FindOne("user", filter, findOptions)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "user_code_login_find_one_error", err
	}

	// 后置处理器
	filter, postResult, msg, msgKey, err := postProcessing(params, filter, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	finalResult, msg, msgKey, err = CodeLoginPost(params, filter, postResult, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, msgKey, err
}

func postProcessing(params t.CodeLoginParams, filter bson.M, result any, c *gin.Context) (bson.M, any, string, string, error) {
	msg, msgKey := "user_code_login Post Processing Success", "user_code_login_post_processing_success"
	var err error = nil
	filter, result, msg, msgKey, err = addJwtToken(params, filter, result, c)
	if err != nil {
		return filter, result, msg, msgKey, err
	}
	// {{占位符 postProcessing}}
	return filter, result, msg, msgKey, err
}
func code2login(params t.CodeLoginParams, filter bson.M, result any, c *gin.Context) (bson.M, bson.M, string, string, error) {
	resultMap, ok := result.(bson.M)
	if !ok {
		errMsg := "result 数据转换为bson.M失败"
		dhlog.Error(errMsg)
		return filter, resultMap, errMsg, "user_code_login_code2login_conver2bson.M_failed", errors.New(errMsg)
	}

	// 定义要POST的URL
	urlPrefix := os.Getenv("OAUTH2_URL_PREFIX")
	urlCode2token := urlPrefix + "/oauth2/code2_token"
	urlAt2userinfo := urlPrefix + "/oauth2/oauth2_get_userinfo"
	urlRefreshToken := urlPrefix + "/oauth2/refresh_token"
	dhlog.DebugAny(filter)
	// 定义要发送的数据
	reqTokenParams := map[string]interface{}{
		"code": filter["code"],
	}
	respToken, err := dhhttp.PostJSON2Map(urlCode2token, reqTokenParams)
	if err != nil {
		dhlog.Error(err.Error())
		return filter, resultMap, err.Error(), respToken["msg_key"].(string), err
	}

	if respToken["status"].(string) != "success" {
		dhlog.Error(respToken["msg"].(string))
		return filter, resultMap, respToken["msg"].(string), respToken["msg_key"].(string), fmt.Errorf(respToken["msg"].(string))
	}

	respTokenData := respToken["data"].(map[string]any)
	accessToken := respTokenData["access_token"].(string)
	refreshToken := respTokenData["refresh_token"].(string)

	// 定义要发送的数据
	reqUserinfoParams := map[string]interface{}{
		"access_token": accessToken,
	}
	respUserinfo, err := dhhttp.PutJSON2Map(urlAt2userinfo, reqUserinfoParams)
	dhlog.DebugAny(respUserinfo)
	if err != nil {
		dhlog.Error(err.Error())
		return filter, resultMap, err.Error(), respUserinfo["msg_key"].(string), err
	}

	if respUserinfo["status"].(string) != "success" {
		// access token 失效，尝试用refresh token刷新
		if respUserinfo["msg_key"].(string) == "oauth2_oauth2_get_userinfo_getUserIdByAt_get_userid_failed" {
			reqNewTokenParams := map[string]interface{}{
				"refresh_token": refreshToken,
			}
			respNewToken, err := dhhttp.PostJSON2Map(urlRefreshToken, reqNewTokenParams)
			if err != nil {
				dhlog.Error(err.Error())
				return filter, resultMap, err.Error(), respNewToken["msg_key"].(string), err
			}
			respNewTokenData := respNewToken["data"].(map[string]any)
			accessToken = respNewTokenData["access_token"].(string)
			refreshToken = respNewTokenData["refresh_token"].(string)

			// 定义要发送的数据
			reqUserinfoParams := map[string]interface{}{
				"access_token": accessToken,
			}
			respUserinfo, err = dhhttp.PutJSON2Map(urlAt2userinfo, reqUserinfoParams)
			if err != nil {
				dhlog.Error(err.Error())
				return filter, resultMap, err.Error(), respUserinfo["msg_key"].(string), err
			}
		} else {
			dhlog.Error(respUserinfo["msg"].(string))
			return filter, resultMap, respUserinfo["msg"].(string), respUserinfo["msg_key"].(string), fmt.Errorf(respUserinfo["msg"].(string))
		}
	}

	// 查询用户是否已经存在
	respUserinfoData := respUserinfo["data"].(map[string]any)
	userFilter := bson.M{}
	userFilter["iuc_id"] = respUserinfoData["_id"].(string)
	count, err := mongodb.Count("user", userFilter)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, nil, err.Error(), "user_code_login_code2login_count_in_db_error", err
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
		userDbData, err := mongodb.MapToBsonD(userinfo)
		if err != nil {
			dhlog.Error(err.Error())
			return userFilter, nil, "数据转成bson.D失败: " + err.Error(), "user_code_login_code2login_mapToBsonD_failed", err
		}

		_, err = mongodb.InsertOneBsonD("user", userDbData)

		if err != nil {
			dhlog.Error(err.Error())
			return userFilter, nil, "数据入库失败: " + err.Error(), "user_code_login_code2login_insert_to_db_failed", err
		}
	}

	return userFilter, nil, "code登录成功", "user_code_login_code2login_success", nil
}

func addJwtToken(params t.CodeLoginParams, filter bson.M, result any, c *gin.Context) (bson.M, bson.M, string, string, error) {
	resultMap, ok := result.(bson.M)
	if !ok {
		errMsg := "result 数据转换为bson.M失败"
		dhlog.Error(errMsg)
		return filter, resultMap, errMsg, "user_code_login_addToken_conver2bson.M_failed", errors.New(errMsg)
	}
	var err error
	resultMap["jwt_access_token"], resultMap["jwt_refresh_token"], err = middleware.JWTGenerateToken(config.JwtSecret, config.JwtRefreshSecret, resultMap["_id"].(primitive.ObjectID).Hex(), resultMap["account"].(string), config.JwtExpSec, config.JwtRefreshSec)
	if err != nil {
		errMsg := err.Error()
		dhlog.Error(errMsg)
		return filter, resultMap, errMsg, "user_code_login_addToken_failed", err
	}
	return filter, resultMap, "生成token成功", "user_code_login_addToken_success", nil
}

// {{占位符 processer}}
