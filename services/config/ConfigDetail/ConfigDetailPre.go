/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-17 21:26:24
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-17 22:13:46
 * @FilePath     : /hecos-v2-api/services/config/ConfigDetail/ConfigDetailPre.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package ConfigDetail

import (
	"github.com/gin-gonic/gin"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func ConfigDetailPre(params t.ConfigDetailParams, c *gin.Context) (bson.M, string, string, error) {
	objUserId := utils.ObjectIDFromHex(params.Id)
	filter := bson.M{"_id": objUserId}

	return filter, "", "config_config_detail_find_one_pre_result", nil
}
