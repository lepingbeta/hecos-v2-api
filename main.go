/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-06 05:06:20
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-08 20:18:06
 * @FilePath     : /hecos-v2-api/main.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"tangxiaoer.shop/dahe/hecos-v2-api/config"
	"tangxiaoer.shop/dahe/hecos-v2-api/lib"
	"tangxiaoer.shop/dahe/hecos-v2-api/routes"
)

func main() {
	// 获取单例实例
	db := mongodb.GetInstance()

	// 连接到 MongoDB
	err := db.Connect(config.MongoURI, 10)
	if err != nil {
		dhlog.Info("Failed to connect to MongoDB:", err)
		return
	}

	// 获取 MongoDB 客户端
	client := db.GetClient()
	dhlog.Info("Connected to MongoDB:", client)

	r := gin.Default()

	// r.Use(Cors())

	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool {
		allowOriginList := []string{"http://192.168.31.11:19823", "http://dev-docker.cezhiqiu.cn:19823"}
		return utils.IsElementInSlice(origin, allowOriginList)
	}
	config.AllowHeaders = []string{"*"}
	// // config.AllowAllOrigins = true

	lib.InitValidator(r)
	// r.Use(cors.New(config))
	routes.SetupRouter(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":18281")
}
