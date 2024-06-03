package CodeLogin

import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

func CodeLoginPre(filter bson.M, c *gin.Context) (bson.M, string, string, error) {

	return filter, "", "user_code_login_find_list_pre_result", nil
}
