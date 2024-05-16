/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-15 23:00:13
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-16 15:25:34
 * @FilePath     : /hecos-v2-api/services/config/ConfigList/ConfigListPre.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package ConfigList

import (
	"github.com/gin-gonic/gin"
	dhjson "github.com/lepingbeta/go-common-v2-dh-json"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func ConfigListPre(params t.ConfigListParams, c *gin.Context) (bson.M, string, string, error) {
	filter := bson.M{"is_delete": 0}
	dhlog.Info(dhjson.JsonEncodeIndent(filter))
	projectId := c.Query("project_id")
	if len(projectId) > 0 {
		filter["project_id"] = projectId
	}
	dhlog.Info(projectId)
	dhlog.Info(dhjson.JsonEncodeIndent(filter))
	return filter, "", "config_config_list_find_list_post_result", nil
}
