package ConfigDetail

import (
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func ConfigDetailPre(params t.ConfigDetailParams, c *gin.Context) (bson.M, string, string, error) {
	filter := bson.M{}

	return filter, "", "config_config_detail_find_one_pre_result", nil
}
