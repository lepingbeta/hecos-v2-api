package DeleteProject

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

type DeleteProject struct {
	Params       t.DeleteProjectParams // 入参结构体版 （原始版）
	DataM        bson.M                // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M              // 入参slice版
	Filter       bson.M                // 入参bson.M版 (查询用)
	DataD        bson.D                // 入参bson.D版 (入库用)
	C            *gin.Context
	Result       any
	Msg          string
	MsgKey       string
	Err          error
	DocID        primitive.ObjectID
}

func (p *DeleteProject) DeleteProject() {

	p.SoftDeleteOne()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}

	fmt.Println("Hello, my name is")
}

func (p *DeleteProject) SoftDeleteOne() {
	p.Filter[`_id`] = mongodb.ObjectIDFromHex(p.Filter[`_id`].(string))
	// {{占位符 composition ObjectIDFromHex}}

	p.DataM["is_delete"] = 1

	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`})
	delete(p.DataM, `_id`)
	// {{占位符 composition filer control}}

	_, p.Err = mongodb.UpdateWithUpdateTime("project", "UpdateOne", p.Filter, p.DataM)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("project_delete_project_SoftDeleteOne 软删除失败：" + p.Err.Error())
		p.MsgKey = "project_delete_project_SoftDeleteOne_to_db_failed"
		return
	}

	p.Msg = "删除成功"
	p.Result = bson.M{"_id": p.Filter["_id"]}
}

// {{占位符 composition}}
