package UpdateConfig

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

type UpdateConfig struct {
	Params       t.UpdateConfigParams // 入参结构体版 （原始版）
	DataM        bson.M               // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M             // 入参slice版
	Filter       bson.M               // 入参bson.M版 (查询用)
	DataD        bson.D               // 入参bson.D版 (入库用)
	C            *gin.Context
	Result       any
	Msg          string
	MsgKey       string
	Err          error
	FindOpts     *options.FindOptions
	DocID        primitive.ObjectID
}

func (p *UpdateConfig) UpdateConfig() {

	p.Convert2ObjectId()
	if p.Err != nil {
		return
	}
	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.CheckExists__0()
	if p.Err != nil {
		return
	}
	p.CheckExists__1()
	if p.Err != nil {
		return
	}
	p.CutFilter()
	if p.Err != nil {
		return
	}
	p.UpdateOne()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *UpdateConfig) Convert2ObjectId() {
	if len(p.Filter[`_id`].(string)) > 0 {
		p.Filter[`_id`] = mongodb.ObjectIDFromHex(p.Filter[`_id`].(string))
	}
}

func (p *UpdateConfig) AddDelete() {
	p.Filter[`is_delete`] = 0
}

func (p *UpdateConfig) CheckExists__0() {
	filter := bson.D{
		{Key: `config_name`, Value: bson.D{{Key: `$eq`, Value: p.Filter[`config_name`]}}},
		{Key: `project_id`, Value: bson.D{{Key: `$eq`, Value: p.Filter[`project_id`]}}},
		{Key: `is_delete`, Value: bson.D{{Key: `$eq`, Value: p.Filter[`is_delete`]}}},
		{Key: `_id`, Value: bson.D{{Key: `$ne`, Value: p.Filter[`_id`]}}},
		// {{占位符}}
	}

	count, err := mongodb.Count("config", filter)
	if err != nil {
		p.Err = err
		p.Msg = utils.DebugMsg("UpdateProject mongodb.Count 查询错误：" + err.Error())
		p.MsgKey = "config_update_config_CheckExists__0_query_db_error"
		dhlog.Error(p.Msg)
		return
	}

	if count > 0 {
		p.Err = err
		p.Msg = utils.DebugMsg("config_update_config_CheckExists__0 没通过")
		p.MsgKey = "config_update_config_CheckExists__0_filter_error"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

}

func (p *UpdateConfig) CheckExists__1() {
	filter := bson.D{
		{Key: `codename`, Value: bson.D{{Key: `$eq`, Value: p.Filter[`codename`]}}},
		{Key: `project_id`, Value: bson.D{{Key: `$eq`, Value: p.Filter[`project_id`]}}},
		{Key: `is_delete`, Value: bson.D{{Key: `$eq`, Value: p.Filter[`is_delete`]}}},
		{Key: `_id`, Value: bson.D{{Key: `$ne`, Value: p.Filter[`_id`]}}},
		// {{占位符}}
	}

	count, err := mongodb.Count("config", filter)
	if err != nil {
		p.Err = err
		p.Msg = utils.DebugMsg("UpdateProject mongodb.Count 查询错误：" + err.Error())
		p.MsgKey = "config_update_config_CheckExists__1_query_db_error"
		dhlog.Error(p.Msg)
		return
	}

	if count > 0 {
		p.Err = err
		p.Msg = utils.DebugMsg("config_update_config_CheckExists__1 没通过")
		p.MsgKey = "config_update_config_CheckExists__1_filter_error"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

}

func (p *UpdateConfig) CutFilter() {
	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`, `is_delete`})
	delete(p.DataM, `_id`)
	delete(p.DataM, `is_delete`)
}

func (p *UpdateConfig) UpdateOne() {
	_, p.Err = mongodb.UpdateWithUpdateTime("config", "UpdateOne", p.Filter, p.DataM)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("数据更新失败：" + p.Err.Error())
		p.MsgKey = "config_update_config_UpdateOne_to_db_failed"
		return
	}

	p.Result = p.DataM
}

// {{占位符 composition}}
