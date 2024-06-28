package UpdateConfig

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	// {{占位符 import}}
)

func (p *UpdateConfig) UpdateConfig() {

	p.CutFilter()
	if p.Err != nil {
		return
	}
	p.UpdateOne()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *UpdateConfig) CutFilter() {
	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`, `is_delete`})
	delete(p.DataM, `_id`)
	delete(p.DataM, `is_delete`)
}

func (p *UpdateConfig) UpdateOne() {
	_, p.Err = mongodb.UpdateWithUpdateTime("config", "UpdateOne", p.Filter, p.DataM)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("数据更新失败：" + p.Err.Error())
		p.MsgKey = "config_update_config_UpdateOne_to_db_failed"
		return
	}

	p.Result = p.DataM
}

// {{占位符 composition}}
