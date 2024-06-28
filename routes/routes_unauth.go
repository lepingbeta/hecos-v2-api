/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 21:08:13
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-28 17:35:13
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
	r.GET("/config/config_list", handlers.ConfigListHandler)
	r.PUT("/config/update_config", handlers.UpdateConfigHandler)
	r.GET("/config/config_detail", handlers.ConfigDetailHandler)
	r.DELETE("/config/delete_config", handlers.DeleteConfigHandler)
	r.PUT("/user/code_login", handlers.CodeLoginHandler)
	r.GET("/project/search_project_list", handlers.SearchProjectListHandler)
	// {{占位符}}
	// r.POST("/auth/add_user", handlers.AddUserHandler)
	// return r
}
