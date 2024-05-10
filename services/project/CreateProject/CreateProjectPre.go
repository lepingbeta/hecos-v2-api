package CreateProject

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

func CreateProjectPre(params t.CreateProjectParams) (bson.D, string, string, error) {
	bsonD, _ := utils.Struct2BsonD(params)
	bsonD = append(bsonD, bson.E{Key: "is_delete", Value: 0})
	return bsonD, "", "", nil
}
