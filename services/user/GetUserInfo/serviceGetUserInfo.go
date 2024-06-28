package GetUserInfo

import (
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	// {{占位符 import}}
)

func (p *GetUserInfo) GetUserInfo() {

	p.PrepareCurrentUser()
	if p.Err != nil {
		return
	}
	p.CutFilter()
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

func (p *GetUserInfo) PrepareCurrentUser() {
	uid, _ := p.C.Get("user_id")
	account, _ := p.C.Get("account")
	p.Filter["_id"] = mongodb.ObjectIDFromHex(uid.(string))
	p.Filter["account"] = account
}

func (p *GetUserInfo) CutFilter() {
	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`})
	delete(p.DataM, `_id`)
}

func (p *GetUserInfo) FindFields() {
	fieldList := bson.D{{Key: "_id", Value: 1},
		{Key: "account", Value: 1},
		{Key: "nickname", Value: 1},
		{Key: "roles", Value: 1}}
	p.FindOneOpts.SetProjection(fieldList)
}

func (p *GetUserInfo) GetListOrOne() {
	result, err := mongodb.FindOne("user", p.Filter, p.FindOneOpts)

	if err != nil {
		dhlog.Error(err.Error())
		p.Msg = utils.DebugMsg("user_get_user_info_GetListOrOne FindOne 错误：" + p.Err.Error())
		p.MsgKey = "user_get_user_info_GetListOrOne_failed"
		return
	}

	p.Result = result
}

// {{占位符 composition}}
