package DeleteConfig

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)


func DeleteConfig(params t.DeleteConfigParams, c *gin.Context) (map[string]interface{}, string, string, error) {

	filter, data, params, msg, msgKey, err := DeleteConfigPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.UpdateOneWithUpdateTime("config", "softDelete", filter, data)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "config_delete_config_soft_delete_db_failed", err
	}


	_, finalResult, msg, msgKey, err := DeleteConfigPost(data, params, c, result)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
