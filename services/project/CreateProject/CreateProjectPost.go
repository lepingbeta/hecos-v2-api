package CreateProject

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func CreateProjectPost(user t.CreateProjectParams, insertId primitive.ObjectID) (map[string]interface{}, string, string, error) {

	finalResult := map[string]interface{}{
		"_id": insertId.Hex(),
	}
	return finalResult, "", "", nil
}
