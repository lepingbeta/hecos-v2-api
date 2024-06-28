package DeleteProject

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	// {{占位符 import}}
)

func (p *DeleteProject) DeleteProject() {

	p.SoftDeleteOne()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *DeleteProject) SoftDeleteOne() {
	p.DataM["is_delete"] = 1

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
