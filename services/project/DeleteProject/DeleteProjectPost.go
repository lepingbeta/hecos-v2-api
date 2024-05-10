package DeleteProject

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"go.mongodb.org/mongo-driver/mongo"
)


// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func DeleteProjectPost(data bson.M, user t.DeleteProjectParams, c *gin.Context, result *mongo.UpdateResult) (bson.M, map[string]interface{}, string, string, error) {

	finalResult := map[string]interface{}{
	}
	return data, finalResult, "", "", nil
}
