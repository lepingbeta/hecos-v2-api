package Test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	"go.mongodb.org/mongo-driver/bson"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
	// {{占位符 import}}
)

type Test struct {
	Params t.EmptyParams
	DataM  bson.M
	Filter bson.M
	DataD  bson.D
	C      *gin.Context
	Result any
	Msg    string
	MsgKey string
	Err    error
}

func (p *Test) Test() {
	p.Insert()
	// {{占位符 composition caller}}

	fmt.Println("Hello, my name is")
}

func (p *Test) Insert() {
	bsonD, _ := mongodb.MapToBsonD(p.DataM)

	p.Result, p.Err = mongodb.InsertOneBsonD("test", bsonD)

	if p.Err != nil {
		dhlog.Error(p.Err.Error())
		p.Msg = "数据入库失败"
		p.MsgKey = "test_test_Insert_to_db_failed"
	}
	fmt.Println("Insert")
	// result, err := mongodb.InsertOneBsonD("test", data)
}

// {{占位符 composition}}
