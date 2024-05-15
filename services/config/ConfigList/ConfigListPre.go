package ConfigList

import (
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func ConfigListPre(params t.ConfigListParams, c *gin.Context) (bson.M, string, string, error) {
	filter := bson.M{"is_delete": 0}

	return filter, "", "config_config_list_find_list_post_result", nil
}
