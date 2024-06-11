package CreateConfig

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

type CreateConfig struct {
	Params       t.CreateConfigParams // 入参结构体版 （原始版）
	DataM        bson.M               // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M             // 入参slice版
	Filter       bson.M               // 入参bson.M版 (查询用)
	DataD        bson.D               // 入参bson.D版 (入库用)
	C            *gin.Context
	Result       any
	Msg          string
	MsgKey       string
	Err          error
	DocID        primitive.ObjectID
}

func (p *CreateConfig) CreateConfig() {

	p.CheckExists()
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

	fmt.Println("Hello, my name is")
}

func (p *CreateConfig) CheckExists() {
	filter := bson.D{
		{Key: `config_name`, Value: bson.D{{Key: `$eq`, Value: p.DataM[`config_name`]}}},
		{Key: `project_id`, Value: bson.D{{Key: `$ne`, Value: p.DataM[`project_id`]}}},
		// {{占位符}}
	}

	count, err := mongodb.Count("config", filter)
	if err != nil {
		p.Err = err
		p.Msg = utils.DebugMsg("UpdateProject mongodb.Count 查询错误：" + err.Error())
		p.MsgKey = "config_create_config_CheckExists_query_db_error"
		dhlog.Error(p.Msg)
		return
	}

	if count > 0 {
		p.Err = err
		p.Msg = utils.DebugMsg("config_create_config_CheckExists 没通过")
		p.MsgKey = "config_create_config_CheckExists_filter_error"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

}

func (p *CreateConfig) AddDelete() {
	p.DataM[`is_delete`] = 0
}

func (p *CreateConfig) Insert() {

	bsonD, _ := mongodb.MapToBsonD(p.DataM)

	var result *mongo.InsertOneResult
	result, p.Err = mongodb.InsertOneBsonD("config", bsonD)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = utils.DebugMsg("数据入库失败：" + p.Err.Error())
		p.MsgKey = "config_create_config_Insert_to_db_failed"
		return
	}

	// 获取并打印 _id
	// 获取并尝试将 _id 转换为 primitive.ObjectID
	var ok bool
	p.DocID, ok = result.InsertedID.(primitive.ObjectID)
	if !ok {
		p.Msg = utils.DebugMsg("Expected the inserted document ID to be a primitive.ObjectID")
		p.MsgKey = "config_create_config_Insert_get_insert_id_failed"
		p.Err = fmt.Errorf(p.Msg)
		dhlog.Error(p.Msg)
		return
	}

	p.Result.(bson.M)["_id"] = p.DocID
}

// {{占位符 composition}}
