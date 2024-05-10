package DeleteProject

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)


func DeleteProject(params t.DeleteProjectParams, c *gin.Context) (map[string]interface{}, string, string, error) {

	filter, data, params, msg, msgKey, err := DeleteProjectPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.UpdateOneWithUpdateTime("project", "softDelete", filter, data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "project_delete_project_insert_to_db_failed", err
	}


	_, finalResult, msg, msgKey, err := DeleteProjectPost(data, params, c, result)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
