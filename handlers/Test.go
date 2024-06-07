package handlers

import (
	"net/http"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	"github.com/lepingbeta/go-common-v2-dh-http/types"
	"tangxiaoer.shop/dahe/hecos-v2-api/services/test/Test"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	dhvalidator "github.com/lepingbeta/go-common-v2-dh-validator"
)

func TestHandler(c *gin.Context) {
	// 处理登录逻辑
	// 声明一个变量来存储 JSON 数据
	var form t.EmptyParams

	respData := types.ResponseData{
		Status: types.ResponseStatus.Success,
		Msg:    "添加成功",
		// MsgKey: "admin_add_user_success",
		MsgKey: "test_test_success",
		Data:   map[string]interface{}{},
	}

	if unsafe.Sizeof(form) > 0 {
		// 使用 BindJSON 方法将 JSON 数据绑定到结构体中
		if err := c.ShouldBindJSON(&form); err != nil {
			respData = types.ResponseData{
				Status: types.ResponseStatus.Error,
				Msg:    err.Error(),
				// MsgKey: "admin_add_user_bind_json_error",
				MsgKey: "test_test_params_error",
				Data:   nil,
			}

			dhlog.Error("参数错误")

			// 如果绑定失败，返回错误信息
			c.JSON(http.StatusBadRequest, respData)
			return
		}

		v, ok := c.MustGet("validator").(*validator.Validate)
		if !ok {
			respData = types.ResponseData{
				Status: types.ResponseStatus.Error,
				Msg:    "Cannot get global validator",
				// MsgKey: "admin_add_user_invalid_validator",
				MsgKey: "test_test_invalid_validator",
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
	}

	dataM, _ := mongodb.Struct2BsonM(form)
	pointer := Test.Test{Params: form, C: c, DataM: dataM, Filter: dataM}
	pointer.Test()
	data := pointer.Result
	msg := pointer.Msg
	msgKey := pointer.MsgKey
	err := pointer.Err

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

	// 返回响应
	c.JSON(http.StatusOK, respData)
}
