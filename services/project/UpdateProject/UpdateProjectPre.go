/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-15 11:19:31
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-15 12:20:57
 * @FilePath     : /hecos-v2-api/services/project/UpdateProject/UpdateProjectPre.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package UpdateProject

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func UpdateProjectPre(data t.UpdateProjectParams, c *gin.Context) (bson.M, t.UpdateProjectParams, string, string, error) {
	objUserId, _ := primitive.ObjectIDFromHex(data.ProjectId)
	filter := bson.M{"_id": objUserId}

	return filter, data, "", "project_update_project_update_pre_result", nil
}
