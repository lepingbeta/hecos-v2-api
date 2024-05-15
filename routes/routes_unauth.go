/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 21:08:13
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-15 19:56:11
 * @FilePath     : /hecos-v2-api/routes/routes_unauth.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package routes

import (
	"github.com/gin-gonic/gin"
	"tangxiaoer.shop/dahe/hecos-v2-api/handlers"
)

func setupUnauthRouter(r *gin.Engine) {
	// r := gin.Default()

	// r.POST("/admin/add_user", handlers.AddUserHandler)
	r.POST("/project/create_project", handlers.CreateProjectHandler)
	r.GET("/project/project_list", handlers.ProjectListHandler)

	r.DELETE("/project/delete_project", handlers.DeleteProjectHandler)
	r.PUT("/project/update_project", handlers.UpdateProjectHandler)
	r.POST("/config/create_config", handlers.CreateConfigHandler)
	// {{占位符}}
	// r.POST("/auth/add_user", handlers.AddUserHandler)
	// return r
}
