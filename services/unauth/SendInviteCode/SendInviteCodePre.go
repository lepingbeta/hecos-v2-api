package SendInviteCode

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func SendInviteCodePre(params t.SendInviteCodeParams, c *gin.Context) (bson.M, string, string, error) {
	filter := bson.M{}

	return filter, "", "", nil
}
