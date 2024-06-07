/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-10 14:50:34
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-05 17:05:49
 * @FilePath     : /hecos-v2-api/services/project/DeleteProject/DeleteProject.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package DeleteProject

import (
	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func DeleteProject(params t.DeleteProjectParams, c *gin.Context) (map[string]interface{}, string, string, error) {

	filter, data, params, msg, msgKey, err := DeleteProjectPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.UpdateWithUpdateTime("project", "softDelete", filter, data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "project_delete_project_insert_to_db_failed", err
	}

	_, finalResult, msg, msgKey, err := DeleteProjectPost(data, params, c, result)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
