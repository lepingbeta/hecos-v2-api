package SendInviteCode

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func SendInviteCodePost(params t.SendInviteCodeParams, result bson.M, c *gin.Context) (map[string]interface{}, string, string, error) {
	// retData := map[string]interface{}{
	// }

	return result, "查询成功", "unauth_send_invite_code_success", nil
}
