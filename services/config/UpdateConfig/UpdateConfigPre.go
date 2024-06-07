/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-16 21:19:54
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-04 22:29:54
 * @FilePath     : /hecos-v2-api/services/config/UpdateConfig/UpdateConfigPre.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package UpdateConfig

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

func UpdateConfigPre(params t.UpdateConfigParams, c *gin.Context) (bson.M, bson.D, string, string, error) {
	objUserId := mongodb.ObjectIDFromHex(params.Id)
	filter := bson.M{"_id": objUserId}
	data, _ := mongodb.Struct2BsonD(params)

	var newDoc bson.D
	for _, elem := range data {
		if elem.Key != "_id" {
			newDoc = append(newDoc, elem)
		}
	}

	return filter, newDoc, "", "config_update_config_update_pre_result", nil
}
