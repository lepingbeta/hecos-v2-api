package CreateConfig

import (
	"fmt"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// {{占位符 import}}
)

func (p *CreateConfig) CreateConfig() {

	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.Insert()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
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
