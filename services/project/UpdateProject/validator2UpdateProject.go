package UpdateProject

import (
	"fmt"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	// {{占位符 import}}
)

func (p *UpdateProject) UpdateProjectValidator2() {
	dhlog.Info("UpdateProjectValidator2")

	p.CheckExists()
	if p.Err != nil {
		return
	}
	// {{占位符 validator2 caller}}
}

func (p *UpdateProject) CheckExists() {
	filter := bson.D{
		{Key: `project_name`, Value: bson.D{{Key: `$eq`, Value: p.Filter[`project_name`]}}},
		{Key: `_id`, Value: bson.D{{Key: `$ne`, Value: mongodb.ObjectIDFromHex(p.Filter[`_id`].(string))}}},
		// {{占位符}}
	}

	count, err := mongodb.Count("project", filter)
	if err != nil {
		p.Err = err
		p.Msg = utils.DebugMsg("UpdateProject mongodb.Count 查询错误：" + err.Error())
		p.MsgKey = "project_update_project_CheckExists_query_db_error"
		dhlog.Error(p.Msg)
		return
	}

	if count > 0 {
		p.Err = err
		p.Msg = utils.DebugMsg("project_update_project_CheckExists 没通过")
		p.MsgKey = "project_update_project_CheckExists_filter_error"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

}

// {{占位符 validator2}}
