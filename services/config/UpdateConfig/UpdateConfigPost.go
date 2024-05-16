package UpdateConfig

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func UpdateConfigPost(params bson.D, c *gin.Context, result *mongo.UpdateResult) (map[string]interface{}, string, string, error) {

	finalResult := map[string]interface{}{}
	return finalResult, "", "config_update_config_update_post_result", nil
}
