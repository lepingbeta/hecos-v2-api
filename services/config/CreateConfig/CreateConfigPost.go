package CreateConfig

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func CreateConfigPost(data bson.D, insertId primitive.ObjectID) (map[string]interface{}, string, string, error) {

	finalResult := map[string]interface{}{
		"_id": insertId.Hex(),
	}
	return finalResult, "", "config_create_config_insert_post_result", nil
}