package SearchProjectList

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

type SearchProjectList struct {
	Params       t.SearchProjectListParams // 入参结构体版 （原始版）
	DataM        bson.M                    // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M                  // 入参slice版
	Filter       bson.M                    // 入参bson.M版 (查询用)
	DataD        bson.D                    // 入参bson.D版 (入库用)
	C            *gin.Context
	Result       any
	Msg          string
	MsgKey       string
	Err          error
	FindOpts     *options.FindOptions
	DocID        primitive.ObjectID
}

func (p *SearchProjectList) SearchProjectList() {

	p.Convert2ObjectId()
	if p.Err != nil {
		return
	}
	p.CutFilter()
	if p.Err != nil {
		return
	}
	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.RemovefEmpty()
	if p.Err != nil {
		return
	}
	p.FindFields()
	if p.Err != nil {
		return
	}
	p.GetList()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *SearchProjectList) Convert2ObjectId() {
	if len(p.Filter[`_id`].(string)) > 0 {
		p.Filter[`_id`] = mongodb.ObjectIDFromHex(p.Filter[`_id`].(string))
	}
}

func (p *SearchProjectList) CutFilter() {
	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`})
	delete(p.DataM, `_id`)
}

func (p *SearchProjectList) AddDelete() {
	p.Filter[`is_delete`] = 0
}

func (p *SearchProjectList) RemovefEmpty() {
	switch v := p.Filter[`_id`].(type) {
	case string:
		if v == "" {
			delete(p.Filter, `_id`)
		}
	case primitive.ObjectID:
		if v.IsZero() {
			delete(p.Filter, `_id`)
		}
	}
}

func (p *SearchProjectList) FindFields() {
	fieldList := bson.D{{Key: "_id", Value: 1},
		{Key: "project_name", Value: 1},
		{Key: "update_callback", Value: 1},
		{Key: "accessId", Value: 1},
		{Key: "create_time", Value: 1}}
	p.FindOpts.SetProjection(fieldList)
}

func (p *SearchProjectList) GetList() {
	result, err := mongodb.FindList("project", p.Filter, p.FindOpts)

	if err != nil {
		dhlog.Error(err.Error())
		p.Msg = utils.DebugMsg("project_search_project_list_GetList FindList 错误：" + p.Err.Error())
		p.MsgKey = "project_search_project_list_GetList_failed"
		return
	}

	p.Result = result
}

// {{占位符 composition}}
