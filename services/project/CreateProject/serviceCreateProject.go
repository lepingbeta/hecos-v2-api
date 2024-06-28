package CreateProject

import (
	"fmt"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	// {{占位符 import}}
)

func (p *CreateProject) CreateProject() {

	p.AddAccessIdAndSecret()
	if p.Err != nil {
		return
	}
	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.Insert()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *CreateProject) AddAccessIdAndSecret() {
	var err error
	p.Result.(bson.M)["accessId"], err = utils.GenerateAccessID(16)
	if err != nil {
		p.Msg = err.Error()
		p.MsgKey = "project_create_project_AddAccessIdAndSecretGenerateAccessID_error"
		dhlog.Error(err.Error())
		return
	}
	p.Result.(bson.M)["accessSecret"], err = utils.GenerateAccessSecret()
	if err != nil {
		p.Msg = err.Error()
		p.MsgKey = "project_create_project_AddAccessIdAndSecretGenerateAccessSecret_error"
		dhlog.Error(err.Error())
		return
	}

	p.DataM["accessId"] = p.Result.(bson.M)["accessId"]

	// 生成随机盐并散列密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Result.(bson.M)["accessSecret"].(string)), bcrypt.DefaultCost)
	if err != nil {
		p.Msg = err.Error()
		p.MsgKey = "project_create_project_AddAccessIdAndSecrethash_password_failed"
		dhlog.Error(err.Error())
		return
	}

	p.DataM["accessSecret"] = string(hashedPassword)
	// result, err := mongodb.InsertOneBsonD("project", data)
}

func (p *CreateProject) AddDelete() {
	p.DataM[`is_delete`] = 0
}

func (p *CreateProject) Insert() {
	bsonD, _ := mongodb.MapToBsonD(p.DataM)

	var result *mongo.InsertOneResult
	result, p.Err = mongodb.InsertOneBsonD("project", bsonD)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("数据入库失败：" + p.Err.Error())
		p.MsgKey = "project_create_project_Insert_to_db_failed"
		return
	}

	// 获取并打印 _id
	// 获取并尝试将 _id 转换为 primitive.ObjectID
	var ok bool
	p.DocID, ok = result.InsertedID.(primitive.ObjectID)
	if !ok {
		p.Msg = utils.DebugMsg("Expected the inserted document ID to be a primitive.ObjectID")
		p.MsgKey = "project_create_project_Insert_get_insert_id_failed"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

	p.Result.(bson.M)["_id"] = p.DocID
}

// {{占位符 composition}}
