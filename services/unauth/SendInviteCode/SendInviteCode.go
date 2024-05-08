/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-06 05:06:20
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-08 20:22:49
 * @FilePath     : /hecos-v2-api/services/unauth/SendInviteCode/SendInviteCode.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package SendInviteCode

import (
	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func SendInviteCode(params t.SendInviteCodeParams, c *gin.Context) (map[string]interface{}, string, string, error) {

	result, msg, msgKey, err := SendInviteCodePre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "unauth_send_invite_code_find_one_error", err
	}

	finalResult, msg, msgKey, err := SendInviteCodePost(params, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, msgKey, err
}
