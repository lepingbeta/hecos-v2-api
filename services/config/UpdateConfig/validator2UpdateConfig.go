package UpdateConfig

import (
	"fmt"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	// {{占位符 import}}
)

func (p *UpdateConfig) UpdateConfigValidator2() {
	dhlog.Info("UpdateConfigValidator2")

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
	// {{占位符 validator2 caller}}
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

// {{占位符 validator2}}
