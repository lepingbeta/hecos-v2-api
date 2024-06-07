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
	Result       bson.M
	Msg          string
	MsgKey       string
	Err          error
	DocID        primitive.ObjectID
}

func (p *CreateProject) CreateProject() {
	p.AddAccessIdAndSecret()
	p.Insert()
	// {{占位符 composition caller}}

	fmt.Println("Hello, my name is")
}

func (p *CreateProject) AddAccessIdAndSecret() {

	var err error
	p.Result["accessId"], err = utils.GenerateAccessID(16)
	if err != nil {
		p.Msg = err.Error()
		p.MsgKey = "project_create_project_AddAccessIdAndSecret_GenerateAccessID_error"
		dhlog.Error(err.Error())
		return
	}
	p.Result["accessSecret"], err = utils.GenerateAccessSecret()
	if err != nil {
		p.Msg = err.Error()
		p.MsgKey = "project_create_project_AddAccessIdAndSecret_GenerateAccessSecret_error"
		dhlog.Error(err.Error())
		return
	}

	p.DataM["accessId"] = p.Result["accessId"]

	// 生成随机盐并散列密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Result["accessSecret"].(string)), bcrypt.DefaultCost)
	if err != nil {
		p.Msg = err.Error()
		p.MsgKey = "project_create_project_AddAccessIdAndSecret_hash_password_failed"
		dhlog.Error(err.Error())
		return
	}

	p.DataM["accessSecret"] = string(hashedPassword)
	// result, err := mongodb.InsertOneBsonD("project", data)
}

func (p *CreateProject) Insert() {
	if true {
		p.DataM["is_delete"] = 0
	}
	bsonD, _ := mongodb.MapToBsonD(p.DataM)

	var result *mongo.InsertOneResult
	result, p.Err = mongodb.InsertOneBsonD("project", bsonD)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = "数据入库失败"
		p.MsgKey = "project_create_project_Insert_to_db_failed"
		return
	}

	// 获取并打印 _id
	// 获取并尝试将 _id 转换为 primitive.ObjectID
	var ok bool
	p.DocID, ok = result.InsertedID.(primitive.ObjectID)
	if !ok {
		p.Msg = "Expected the inserted document ID to be a primitive.ObjectID"
		p.MsgKey = "project_create_project_Insert_get_insert_id_failed"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

	p.Result["_id"] = p.DocID
}

// {{占位符 composition}}
