/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-15 20:05:45
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-02 14:13:10
 * @FilePath     : /hecos-v2-api/services/config/CreateConfig/CreateConfig.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package CreateConfig

import (
	"errors"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func preProcessing(params t.CreateConfigParams) (map[string]interface{}, string, string, error) {
	filter := bson.D{
		{Key: `config_name`, Value: bson.D{{Key: `$eq`, Value: params.ConfigName}}},
		{Key: `project_id`, Value: bson.D{{Key: `$eq`, Value: params.ProjectId}}},
		// {{占位符}}
	}

	count, err := mongodb.Count("config", filter)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, "CreateConfig mongodb.Count 查询错误：" + err.Error(), "CreateConfig_query_db_error", err
	}

	if count > 0 {
		errMsg := "CreateConfig filter没通过"
		err = errors.New(errMsg)
		dhlog.Error(errMsg)
		return nil, errMsg, "CreateConfig_msg_key_filter_error", err
	}

	// return nil, "", "", nil
	// {{占位符 preProcessing}}
	return nil, "", "", nil
}

func CreateConfig(params t.CreateConfigParams) (map[string]interface{}, string, string, error) {
	finalResult, msg, msgKey, err := preProcessing(params)
	if err != nil {
		return finalResult, msg, msgKey, err
	}

	data, msg, msgKey, err := CreateConfigPre(params)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.InsertOneBsonD("config", data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, "数据入库失败", "config_create_config_insert_to_db_failed", err
	}

	// 获取并打印 _id
	// 获取并尝试将 _id 转换为 primitive.ObjectID
	docID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		dhlog.Error("Expected the inserted document ID to be a primitive.ObjectID")
		return nil, "Expected the inserted document ID to be a primitive.ObjectID", "config_create_config_insert_id_error", err
	}

	finalResult, msg, msgKey, err = CreateConfigPost(data, docID)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
