package ConfigDetail

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	"github.com/gin-gonic/gin"
)

func ConfigDetail(params t.ConfigDetailParams, c *gin.Context) (map[string]interface{}, string, string, error) {

	filter, msg, msgKey, err := ConfigDetailPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.FindOne("config", filter)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "config_config_detail_find_one_error", err
	}

	finalResult, msg, msgKey, err := ConfigDetailPost(params, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, msgKey, err
}
