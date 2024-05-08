package routes

import (
	"github.com/gin-gonic/gin"
	"tangxiaoer.shop/dahe/hecos-v2-api/handlers"
)

func setupUnauthRouter(r *gin.Engine) {
	// r := gin.Default()

	// r.POST("/admin/add_user", handlers.AddUserHandler)
	r.GET("/unauth/send_invite_code", handlers.SendInviteCodeHandler)
	r.POST("/unauth/create_project", handlers.CreateProjectHandler)
	// {{占位符}}
	// r.POST("/auth/add_user", handlers.AddUserHandler)
	// return r
}
