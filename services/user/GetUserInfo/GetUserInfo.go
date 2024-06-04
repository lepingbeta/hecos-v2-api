package GetUserInfo

import (
	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

func preProcessing(params t.EmptyParams, filterEmpty bson.M, resultEmpty any, c *gin.Context) (bson.M, any, string, string, error) {
	filter, _ := utils.Struct2BsonM(params)
	var result = bson.M{}
	var msg, msgKey string = "", ""
	var err error = nil

	// {{占位符 preProcessing}}

	return filter, result, msg, msgKey, err
}

func GetUserInfo(params t.EmptyParams, c *gin.Context) (any, string, string, error) {
	filter, finalResult, msg, msgKey, err := preProcessing(params, nil, nil, c)
	// {{bsonD holder}}
	if err != nil {
		return finalResult, msg, msgKey, err
	}

	filter, msg, msgKey, err = GetUserInfoPre(filter, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	fieldList := bson.D{
		{Key: "_id", Value: 1},
		{Key: "account", Value: 1},
		{Key: "nickname", Value: 1},
		{Key: "roles", Value: 1},
	}
	// 创建Find选项，设置Projection
	findOptions := options.FindOne()
	findOptions.SetProjection(fieldList)
	result, err := mongodb.FindOne("user", filter, findOptions)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "user_get_user_info_find_one_error", err
	}

	// 后置处理器
	filter, postResult, msg, msgKey, err := postProcessing(params, filter, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	finalResult, msg, msgKey, err = GetUserInfoPost(params, filter, postResult, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, msgKey, err
}

func postProcessing(params t.EmptyParams, filter bson.M, result any, c *gin.Context) (bson.M, any, string, string, error) {
	msg, msgKey := "user_get_user_info Post Processing Success", "user_get_user_info_post_processing_success"
	var err error = nil
	// {{占位符 postProcessing}}
	return filter, result, msg, msgKey, err
}

// {{占位符 processer}}
