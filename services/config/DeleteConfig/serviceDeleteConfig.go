package DeleteConfig

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	// {{占位符 import}}
)

func (p *DeleteConfig) DeleteConfig() {

	p.Convert2ObjectId()
	if p.Err != nil {
		return
	}
	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.CutFilter()
	if p.Err != nil {
		return
	}
	p.SoftDeleteOne()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *DeleteConfig) Convert2ObjectId() {
	if len(p.Filter[`_id`].(string)) > 0 {
		p.Filter[`_id`] = mongodb.ObjectIDFromHex(p.Filter[`_id`].(string))
	}
}

func (p *DeleteConfig) AddDelete() {
	p.Filter[`is_delete`] = 0
}

func (p *DeleteConfig) CutFilter() {
	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`, `is_delete`})
	delete(p.DataM, `_id`)
	delete(p.DataM, `is_delete`)
}

func (p *DeleteConfig) SoftDeleteOne() {
	p.DataM["is_delete"] = 1

	_, p.Err = mongodb.UpdateWithUpdateTime("config", "UpdateOne", p.Filter, p.DataM)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("config_delete_config_SoftDeleteOne 软删除失败：" + p.Err.Error())
		p.MsgKey = "config_delete_config_SoftDeleteOne_to_db_failed"
		return
	}

	p.Msg = "删除成功"
	p.Result = bson.M{"_id": p.Filter["_id"]}
}

// {{占位符 composition}}
