package ProjectList

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ProjectList(params t.ProjectListParams, c *gin.Context) ([]bson.M, string, string, error) {

	filter, msg, msgKey, err := ProjectListPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.FindList("project", filter)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "unauth_project_list_find_one_error", err
	}

	finalResult, msg, msgKey, err := ProjectListPost(params, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, msgKey, err
}
