package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validatorMiddleware(v *validator.Validate) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("validator", v)
		c.Next()
	}
}

// setupValidator 初始化并注册自定义验证器，返回一个配置好的validator实例
func setupValidator() *validator.Validate {
	validate := validator.New()
	loadCommonValidator(validate)
	loadManualValidator(validate)
	return validate
}

func InitValidator(r *gin.Engine) {
	validate := setupValidator()
	r.Use(validatorMiddleware(validate))

}
