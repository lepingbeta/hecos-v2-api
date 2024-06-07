package UpdateProject

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

type UpdateProject struct {
	Params       t.UpdateProjectParams // 入参结构体版 （原始版）
	DataM        bson.M                // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M              // 入参slice版
	Filter       bson.M                // 入参bson.M版 (查询用)
	DataD        bson.D                // 入参bson.D版 (入库用)
	C            *gin.Context
	Result       bson.M
	Msg          string
	MsgKey       string
	Err          error
	DocID        primitive.ObjectID
}

func (p *UpdateProject) UpdateProject() {

	p.CheckExists()
	if p.Err != nil {
		return
	}
	p.UpdateOne()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}

	fmt.Println("Hello, my name is")
}

func (p *UpdateProject) CheckExists() {
	filter := bson.D{
		{Key: `project_name`, Value: bson.D{{Key: `$eq`, Value: p.DataM[`project_name`]}}},
		{Key: `_id`, Value: bson.D{{Key: `$ne`, Value: mongodb.ObjectIDFromHex(p.DataM[`_id`].(string))}}},
		// {{占位符}}
	}

	count, err := mongodb.Count("project", filter)
	if err != nil {
		p.Err = err
		p.Msg = "UpdateProject mongodb.Count 查询错误：" + err.Error()
		p.MsgKey = "project_update_project_CheckExists_query_db_error"
		dhlog.Error(p.Msg)
		return
	}

	if count > 0 {
		p.Err = err
		p.Msg = "project_update_project_CheckExists 没通过"
		p.MsgKey = "project_update_project_CheckExists_filter_error"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

}

func (p *UpdateProject) UpdateOne() {
	p.Filter[`_id`] = mongodb.ObjectIDFromHex(p.Filter[`_id`].(string))
	// {{占位符 composition ObjectIDFromHex}}

	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`})
	delete(p.DataM, `_id`)
	// {{占位符 composition filer control}}

	_, p.Err = mongodb.UpdateWithUpdateTime("project", "UpdateOne", p.Filter, p.DataM)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = "数据更新失败"
		p.MsgKey = "project_update_project_UpdateOne_to_db_failed"
		return
	}

	p.Result = p.DataM
}

// {{占位符 composition}}
