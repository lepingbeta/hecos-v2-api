package UpdateProject

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	// {{占位符 import}}
)

func (p *UpdateProject) UpdateProject() {

	p.UpdateOne()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *UpdateProject) UpdateOne() {
	_, p.Err = mongodb.UpdateWithUpdateTime("project", "UpdateOne", p.Filter, p.DataM)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("数据更新失败：" + p.Err.Error())
		p.MsgKey = "project_update_project_UpdateOne_to_db_failed"
		return
	}

	p.Result = p.DataM
}

// {{占位符 composition}}
