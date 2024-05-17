package ConfigDetail

import (

	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func ConfigDetailPost(params t.ConfigDetailParams, result bson.M, c *gin.Context) (map[string]interface{}, string, string, error) {
	// retData := map[string]interface{}{
	// }

	return result, "查询成功", "config_config_detail_find_one_post_result", nil
}
