package ConfigList

import (
	"github.com/gin-gonic/gin"
	httpTypes "github.com/lepingbeta/go-common-v2-dh-http/types"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
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

func (p *ConfigList) ConfigList() {

	p.CutFilter()
	if p.Err != nil {
		return
	}
	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.RemovefEmpty()
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
	}
	p.TempPrepare()
	if p.Err != nil {
		return
	}
	p.QueryOtherCollection()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *ConfigList) CutFilter() {
	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`project_id`})
	delete(p.DataM, `project_id`)
}

func (p *ConfigList) AddDelete() {
	p.Filter[`is_delete`] = 0
}

func (p *ConfigList) RemovefEmpty() {
	switch v := p.Filter[`project_id`].(type) {
	case string:
		if v == "" {
			delete(p.Filter, `project_id`)
		}
	case primitive.ObjectID:
		if v.IsZero() {
			delete(p.Filter, `project_id`)
		}
	}
}

func (p *ConfigList) FindFields() {
	fieldList := bson.D{{Key: "_id", Value: 1},
		{Key: "project_id", Value: 1},
		{Key: "config_name", Value: 1},
		{Key: "codename", Value: 1},
		{Key: "use_callback", Value: 1},
		{Key: "create_time", Value: 1},
		{Key: "update_time", Value: 1}}
	p.FindOpts.SetProjection(fieldList)
}

func (p *ConfigList) GetListOrOne() {
	result, err := mongodb.FindList("config", p.Filter, p.FindOpts)

	if err != nil {
		dhlog.Error(err.Error())
		p.Msg = utils.DebugMsg("config_config_list_GetListOrOne FindList 错误：" + p.Err.Error())
		p.MsgKey = "config_config_list_GetListOrOne_failed"
		return
	}

	p.Result = result
}

func (p *ConfigList) TempPrepare() {
	method := "slice" // slice,pager
	switch method {
	case "bson.M":
		// 将 bson.M 对象放入切片
		p.Temp1 = append(p.Temp1, p.Result.(bson.M))
	case "slice":
		// 将 bson.M 对象放入切片
		p.Temp1 = p.Result.([]bson.M)
	case "pager":
		p.Temp1 = p.Result.(httpTypes.DataList).List
	}
}

func (p *ConfigList) QueryOtherCollection() {

	var sourceValueList []primitive.ObjectID
	// 遍历切片
	for _, item := range p.Temp1 {
		sourceValue := mongodb.ObjectIDFromHex(item[`project_id`].(string))

		sourceValueList = append(sourceValueList, sourceValue)
	}
	filter2 := bson.M{"_id": bson.M{"$in": sourceValueList}}
	fieldList := bson.D{{Key: "project_name", Value: 1}, {Key: "update_callback", Value: 1}, {Key: "_id", Value: 1}}
	// 创建Find选项，设置Projection
	findOptions := options.Find()
	findOptions.SetProjection(fieldList)
	result2, err := mongodb.FindList("project", filter2, findOptions)
	if err != nil {
		return
	}

	// 创建一个映射，以project_id为键，以project_name为值
	idToDetailsMap := make(map[any]bson.M)
	for _, res2 := range result2 {
		keyId := res2["_id"]

		// 复制res2的内容到新的bson.M映射中
		details := bson.M{}
		// 假设我们知道result中包含的字段
		fieldsInResult := []string{"project_name", "update_callback"}
		for _, key := range fieldsInResult {
			if res2[key] != nil { // 检查键是否存在并且有值
				details[key] = res2[key]
			}
		}
		idToDetailsMap[keyId] = details
	}

	for i, item := range p.Temp1 {
		keyId := mongodb.ObjectIDFromHex(item[`project_id`].(string))
		if details, exists := idToDetailsMap[keyId]; exists {
			// 将details中的字段添加到result[i]中
			for key, value := range details {
				// 检查result[i]中是否存在该字段，如果存在则添加
				if _, exists := p.Temp1[i][key]; !exists {
					p.Temp1[i][key] = value
				}
			}
		}
	}

}

// {{占位符 composition}}
