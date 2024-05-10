package DeleteProject

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func DeleteProjectPre(params t.DeleteProjectParams, c *gin.Context) (bson.M, bson.M, t.DeleteProjectParams, string, string, error) {
	id := c.Query("id")
	objUserId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objUserId}

	data := bson.M{"is_delete": 1}
	return filter, data, params, "", "", nil
}
