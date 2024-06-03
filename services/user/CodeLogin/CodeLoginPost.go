package CodeLogin

import (

	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func CodeLoginPost(params t.CodeLoginParams, filter bson.M, result any, c *gin.Context) (any, string, string, error) {
	// retData := map[string]interface{}{
	// }

	return result, "查询成功", "user_code_login_find_list_post_result", nil
}
