package DeleteConfig

import (
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

type DeleteConfig struct {
	Params       t.DeleteConfigParams // 入参结构体版 （原始版）
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
