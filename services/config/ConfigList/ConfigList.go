package ConfigList

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ConfigList(params t.ConfigListParams, c *gin.Context) ([]bson.M, string, string, error) {

	filter, msg, msgKey, err := ConfigListPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.FindList("config", filter)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "config_config_list_find_list_error", err
	}

	finalResult, msg, msgKey, err := ConfigListPost(params, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, msgKey, err
}
