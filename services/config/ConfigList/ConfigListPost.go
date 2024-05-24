/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-15 23:00:13
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-20 16:29:57
 * @FilePath     : /hecos-v2-api/services/config/ConfigList/ConfigListPost.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package ConfigList

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func ConfigListPost(params t.ConfigListParams, filter bson.M, result []bson.M, c *gin.Context) ([]bson.M, string, string, error) {
	// retData := map[string]interface{}{
	// }

	return result, "查询成功", "config_config_list_find_list_post_result", nil
}
