package CreateProject

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

type CreateProject struct {
	Params       t.CreateProjectParams // 入参结构体版 （原始版）
	DataM        bson.M                // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M              // 入参slice版
	Filter       bson.M                // 入参bson.M版 (查询用)
	DataD        bson.D                // 入参bson.D版 (入库用)
	C            *gin.Context
	Result       any
	Msg          string
	MsgKey       string
	Err          error
	FindOpts     *options.FindOptions
	FindOneOpts  *options.FindOneOptions
	DocID        primitive.ObjectID
	// 临时变量3兄弟
	Temp1 []bson.M
	Temp2 any
	Temp3 any
}

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
