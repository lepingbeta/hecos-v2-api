package DeleteConfig

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func DeleteConfigPre(params t.DeleteConfigParams, c *gin.Context) (bson.M, bson.M, t.DeleteConfigParams, string, string, error) {
	objUserId := utils.ObjectIDFromHex(params.Id)
	filter := bson.M{"_id": objUserId}

	data := bson.M{"is_delete": 1}
	return filter, data, params, "", "config_delete_config_soft_delete_pre_result", nil
}
