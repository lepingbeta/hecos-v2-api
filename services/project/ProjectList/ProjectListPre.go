/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-09 18:26:50
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-16 15:25:25
 * @FilePath     : /hecos-v2-api/services/project/ProjectList/ProjectListPre.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package ProjectList

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func ProjectListPre(params t.ProjectListParams, c *gin.Context) (bson.M, string, string, error) {
	filter := bson.M{"is_delete": 0}

	return filter, "", "", nil
}
