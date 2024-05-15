package ConfigList

import (

	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func ConfigListPost(params t.ConfigListParams, result []bson.M, c *gin.Context) ([]bson.M, string, string, error) {
	// retData := map[string]interface{}{
	// }

	return result, "查询成功", "config_config_list_find_list_post_result", nil
}
