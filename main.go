/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-06 05:06:20
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-04 07:53:12
 * @FilePath     : /hecos-v2-api/main.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"tangxiaoer.shop/dahe/hecos-v2-api/lib"
	"tangxiaoer.shop/dahe/hecos-v2-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		dhlog.Error(err.Error())
	}

	// 获取单例实例
	db := mongodb.GetInstance()

	// 连接到 MongoDB
	mongoUri := os.Getenv("MongoURI")
	err = db.Connect(mongoUri, 10)
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
		allowOriginList := []string{"http://192.168.31.11:29517", "http://dev-docker.cezhiqiu.cn:29517", "http://localhost:29517"}
		return utils.IsElementInSlice(origin, allowOriginList)
	}
	config.AllowHeaders = []string{"*"}
	// // config.AllowAllOrigins = true

	lib.InitValidator(r)
	r.Use(cors.New(config))
	routes.SetupRouter(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":18281")
}
