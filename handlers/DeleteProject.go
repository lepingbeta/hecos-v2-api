package handlers

import (
	"net/http"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lepingbeta/go-common-v2-dh-http/types"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	dhvalidator "github.com/lepingbeta/go-common-v2-dh-validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tangxiaoer.shop/dahe/hecos-v2-api/services/project/DeleteProject"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func DeleteProjectHandler(c *gin.Context) {
	// 处理登录逻辑
	// 声明一个变量来存储 JSON 数据
	var form t.DeleteProjectParams

	respData := types.ResponseData{
		Status: types.ResponseStatus.Success,
		Msg:    "成功",
		// MsgKey: "admin_add_user_success",
		MsgKey: "project_delete_project_success",
		Data:   map[string]interface{}{},
	}

	if unsafe.Sizeof(form) > 0 {
		// 使用 BindJSON 方法将 JSON 数据绑定到结构体中
		if err := c.ShouldBindJSON(&form); err != nil {
			respData = types.ResponseData{
				Status: types.ResponseStatus.Error,
				Msg:    err.Error(),
				// MsgKey: "admin_add_user_bind_json_error",
				MsgKey: "project_delete_project_params_error",
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
				MsgKey: "project_delete_project_invalid_validator",
				Data:   nil,
			}

			c.JSON(http.StatusInternalServerError, respData)
			return
		}

		// 对参数进行简单验证
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
	pointer := DeleteProject.DeleteProject{Params: form, C: c, DataM: dataM, Filter: dataM, Result: bson.M{}, FindOpts: options.Find(), FindOneOpts: options.FindOne()}

	// 对参数进行复杂验证
	pointer.DeleteProjectValidator2()
	if pointer.Err != nil {
		respData = types.ResponseData{
			Status: types.ResponseStatus.Error,
			MsgKey: pointer.MsgKey,
			Msg:    pointer.Err.Error() + " || " + pointer.Msg,
			Data:   pointer.Result,
		}
		// 返回响应
		c.JSON(http.StatusOK, respData)
		return
	}

	// 对数据进行处理
	pointer.DeleteProject()
	if pointer.Err != nil {
		respData = types.ResponseData{
			Status: types.ResponseStatus.Error,
			MsgKey: pointer.MsgKey,
			Msg:    pointer.Err.Error() + " || " + pointer.Msg,
			Data:   pointer.Result,
		}
	}

	// 写入成功消息
	if len(pointer.Msg) > 0 {
		respData.Msg = pointer.Msg
	}

	// 写入成功消息key
	if len(pointer.MsgKey) > 0 {
		respData.MsgKey = pointer.MsgKey
	}

	respData.Data = pointer.Result

	// 返回响应
	c.JSON(http.StatusOK, respData)
}
