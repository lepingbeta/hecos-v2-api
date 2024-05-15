package UpdateProject

import (
	"github.com/gin-gonic/gin"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"go.mongodb.org/mongo-driver/mongo"
)


// User 结构体表示用户信息
// type User struct {
// 	Username string
// 	Password string
// }

func UpdateProjectPost(user t.UpdateProjectParams, c *gin.Context, result *mongo.UpdateResult) (map[string]interface{}, string, string, error) {

	finalResult := map[string]interface{}{
	}
	return finalResult, "", "project_update_project_update_post_result", nil
}
