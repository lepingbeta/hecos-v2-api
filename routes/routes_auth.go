/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-05 07:58:26
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-05 17:58:55
 * @FilePath     : /dahe/hecos-v2-api/routes/routes_auth.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package routes

import (
	"github.com/gin-gonic/gin"
	"tangxiaoer.shop/dahe/hecos-v2-api/handlers"
)

func setupAuthRouter(r *gin.Engine) {
	// r := gin.Default()

	r.GET("/user/get_user_info", handlers.GetUserInfoHandler)
	// {{占位符}}
	// return r
}
