package CreateConfig

import (
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func CreateConfigPre(params t.CreateConfigParams) (bson.D, string, string, error) {
	bsonD, _ := utils.Struct2BsonD(params)
	bsonD = append(bsonD, bson.E{Key: "is_delete", Value: 0})
	return bsonD, "", "config_create_config_insert_pre_result", nil
}
