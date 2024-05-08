/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-06 05:06:20
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-08 20:23:17
 * @FilePath     : /hecos-v2-api/routes/routes.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lepingbeta/go-common-v2-dh-middleware"
	"tangxiaoer.shop/dahe/hecos-v2-api/config"
)

func SetupRouter(r *gin.Engine) {
	// r := gin.Default()

	setupUnauthRouter(r)

	// auth := r.Group("/auth")
	r.Use(middleware.JWTParseToken(config.JwtSecret))
	setupAuthRouter(r)
	// return r
}
