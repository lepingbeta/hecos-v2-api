package GetUserInfo

import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

func GetUserInfoPre(filter bson.M, c *gin.Context) (bson.M, string, string, error) {

	return filter, "", "user_get_user_info_find_list_pre_result", nil
}
