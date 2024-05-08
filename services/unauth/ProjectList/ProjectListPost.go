package ProjectList

import (

	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func ProjectListPost(params t.ProjectListParams, result []bson.M, c *gin.Context) ([]bson.M, string, string, error) {
	// retData := map[string]interface{}{
	// }

	return result, "查询成功", "unauth_project_list_success", nil
}
