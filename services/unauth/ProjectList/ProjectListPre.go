package ProjectList

import (
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func ProjectListPre(params t.ProjectListParams, c *gin.Context) (bson.M, string, string, error) {
	filter := bson.M{}

	return filter, "", "", nil
}
