package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lepingbeta/go-common-v2-dh-http/types"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	dhvalidator "github.com/lepingbeta/go-common-v2-dh-validator"
	"tangxiaoer.shop/dahe/hecos-v2-api/services/unauth/SendInviteCode"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func SendInviteCodeHandler(c *gin.Context) {
	// 处理登录逻辑
	// 声明一个变量来存储 JSON 数据
	var form t.SendInviteCodeParams
	respData := types.ResponseData{
		Status: types.ResponseStatus.Success,
		Msg:    "查询成功",
		MsgKey: "unauth_send_invite_code_success",
		Data:   map[string]interface{}{},
	}

	// 使用 BindJSON 方法将 JSON 数据绑定到结构体中
	if err := c.ShouldBindJSON(&form); err != nil {
		respData = types.ResponseData{
			Status: types.ResponseStatus.Error,
			Msg:    err.Error(),
			MsgKey: "unauth_send_invite_code_params_error",
			Data:   nil,
		}
		// 如果绑定失败，返回错误信息
		c.JSON(http.StatusBadRequest, respData)
		return
	}

	v, ok := c.MustGet("validator").(*validator.Validate)
	if !ok {
		respData = types.ResponseData{
			Status: types.ResponseStatus.Error,
			Msg:    "Cannot get global validator",
			MsgKey: "unauth_send_invite_code_invalid_validator",
			Data:   nil,
		}

		c.JSON(http.StatusInternalServerError, respData)
		return
	}

	if err := v.Struct(form); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				dhlog.Error(dhvalidator.CustomErrors(e))

				respData = types.ResponseData{
					Status: types.ResponseStatus.Error,
					Msg:    err.Error(),
					MsgKey: dhvalidator.CustomErrors(e),
					Data:   nil,
				}
				c.JSON(http.StatusInternalServerError, respData)
				return
			}
		}
	}

	data, msg, msgKey, err := SendInviteCode.SendInviteCode(form, c)

	if err != nil {
		respData = types.ResponseData{
			Status: types.ResponseStatus.Error,
			MsgKey: msgKey,
			Msg:    err.Error() + " || " + msg,
			Data:   data,
		}
	}

	// 写入成功消息
	if len(msg) > 0 {
		respData.Msg = msg
	}

	// 写入成功消息key
	if len(msgKey) > 0 {
		respData.MsgKey = msgKey
	}

	respData.Data = data

	c.JSON(http.StatusOK, respData)
}