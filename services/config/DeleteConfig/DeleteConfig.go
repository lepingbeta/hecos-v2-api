/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-18 04:55:31
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-05 17:05:58
 * @FilePath     : /hecos-v2-api/services/config/DeleteConfig/DeleteConfig.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package DeleteConfig

import (
	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func DeleteConfig(params t.DeleteConfigParams, c *gin.Context) (map[string]interface{}, string, string, error) {

	filter, data, params, msg, msgKey, err := DeleteConfigPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.UpdateWithUpdateTime("config", "softDelete", filter, data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "config_delete_config_soft_delete_db_failed", err
	}

	_, finalResult, msg, msgKey, err := DeleteConfigPost(data, params, c, result)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
