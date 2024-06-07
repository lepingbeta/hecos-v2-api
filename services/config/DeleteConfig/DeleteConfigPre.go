/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-18 04:55:31
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-04 22:29:41
 * @FilePath     : /hecos-v2-api/services/config/DeleteConfig/DeleteConfigPre.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package DeleteConfig

import (
	"github.com/gin-gonic/gin"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func DeleteConfigPre(params t.DeleteConfigParams, c *gin.Context) (bson.M, bson.M, t.DeleteConfigParams, string, string, error) {
	objUserId := mongodb.ObjectIDFromHex(params.Id)
	filter := bson.M{"_id": objUserId}

	data := bson.M{"is_delete": 1}
	return filter, data, params, "", "config_delete_config_soft_delete_pre_result", nil
}
