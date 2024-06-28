package ConfigList

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

type ConfigList struct {
	Params       t.ConfigListParams // 入参结构体版 （原始版）
	DataM        bson.M             // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M           // 入参slice版
	Filter       bson.M             // 入参bson.M版 (查询用)
	DataD        bson.D             // 入参bson.D版 (入库用)
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
