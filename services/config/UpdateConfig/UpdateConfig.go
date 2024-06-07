/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-16 21:19:54
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-05 17:05:54
 * @FilePath     : /hecos-v2-api/services/config/UpdateConfig/UpdateConfig.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package UpdateConfig

import (
	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func preProcessing(params t.UpdateConfigParams) (map[string]interface{}, string, string, error) {
	// {{占位符 preProcessing}}
	return nil, "", "", nil
}

func UpdateConfig(params t.UpdateConfigParams, c *gin.Context) (map[string]interface{}, string, string, error) {
	finalResult, msg, msgKey, err := preProcessing(params)
	if err != nil {
		return finalResult, msg, msgKey, err
	}

	filter, data, msg, msgKey, err := UpdateConfigPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.UpdateWithUpdateTime("config", "UpdateOne", filter, data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "config_update_config_update_db_failed", err
	}

	finalResult, msg, msgKey, err = UpdateConfigPost(data, c, result)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
