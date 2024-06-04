package GetUserInfo

import (

	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func GetUserInfoPost(params t.EmptyParams, filter bson.M, result any, c *gin.Context) (any, string, string, error) {
	// retData := map[string]interface{}{
	// }

	return result, "查询成功", "user_get_user_info_find_list_post_result", nil
}
