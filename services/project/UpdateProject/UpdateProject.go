package UpdateProject

import (
	"errors"

	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func preProcessing(params t.UpdateProjectParams) (map[string]interface{}, string, string, error) {
	filter := bson.D{
		{Key: `project_name`, Value: bson.D{{Key: `$eq`, Value: params.ProjectName}}},
		{Key: `_id`, Value: bson.D{{Key: `$ne`, Value: utils.ObjectIDFromHex(params.ProjectId)}}},
		// {{占位符}}
	}

	count, err := mongodb.Count("project", filter)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, "UpdateProject mongodb.Count 查询错误：" + err.Error(), "UpdateProject_query_db_error", err
	}

	if count > 0 {
		errMsg := "UpdateProject filter没通过"
		err = errors.New(errMsg)
		dhlog.Error(errMsg)
		return nil, errMsg, "UpdateProject_msg_key_filter_error", err
	}

	// return nil, "", "", nil
	// {{占位符 preProcessing}}
	return nil, "", "", nil
}

func UpdateProject(data t.UpdateProjectParams, c *gin.Context) (map[string]interface{}, string, string, error) {
	finalResult, msg, msgKey, err := preProcessing(data)
	if err != nil {
		return finalResult, msg, msgKey, err
	}

	filter, data, msg, msgKey, err := UpdateProjectPre(data, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.UpdateOneWithUpdateTime("project", "UpdateOne", filter, data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "project_update_project_update_db_failed", err
	}

	finalResult, msg, msgKey, err = UpdateProjectPost(data, c, result)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
