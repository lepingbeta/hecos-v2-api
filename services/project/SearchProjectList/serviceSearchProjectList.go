package SearchProjectList

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// {{占位符 import}}
)

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
	p.GetListOrOne()
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

func (p *SearchProjectList) GetListOrOne() {
	result, err := mongodb.FindList("project", p.Filter, p.FindOpts)

	if err != nil {
		dhlog.Error(err.Error())
		p.Msg = utils.DebugMsg("project_search_project_list_GetListOrOne FindList 错误：" + p.Err.Error())
		p.MsgKey = "project_search_project_list_GetListOrOne_failed"
		return
	}

	p.Result = result
}

// {{占位符 composition}}
