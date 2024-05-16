package UpdateConfig

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func preProcessing(params t.UpdateConfigParams) (map[string]interface{}, string, string, error) {
	// {{占位符 preProcessing}}
	return nil, "", "", nil
}

func UpdateConfig(params t.UpdateConfigParams, c *gin.Context) (map[string]interface{}, string, string, error) {
	finalResult, msg, msgKey, err := preProcessing(params)
	if err != nil {
		return finalResult, msg, msgKey, err
	}

	filter, data, msg, msgKey, err := UpdateConfigPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.UpdateOneWithUpdateTime("config", "UpdateOne", filter, data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "config_update_config_update_db_failed", err
	}

	finalResult, msg, msgKey, err = UpdateConfigPost(data, c, result)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
