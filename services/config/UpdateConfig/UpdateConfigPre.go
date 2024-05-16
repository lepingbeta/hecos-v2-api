package UpdateConfig

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

func UpdateConfigPre(params t.UpdateConfigParams, c *gin.Context) (bson.M, bson.D, string, string, error) {
	objUserId := utils.ObjectIDFromHex(params.Id)
	filter := bson.M{"_id": objUserId}
	data, _ := utils.Struct2BsonD(params)

	var newDoc bson.D
	for _, elem := range data {
		if elem.Key != "_id" {
			newDoc = append(newDoc, elem)
		}
	}

	return filter, newDoc, "", "config_update_config_update_pre_result", nil
}
