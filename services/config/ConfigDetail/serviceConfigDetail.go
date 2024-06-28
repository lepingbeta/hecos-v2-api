package ConfigDetail

import (
	httpTypes "github.com/lepingbeta/go-common-v2-dh-http/types"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// {{占位符 import}}
)

func (p *ConfigDetail) ConfigDetail() {

	p.CutFilter()
	if p.Err != nil {
		return
	}
	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.Convert2ObjectId()
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

func (p *ConfigDetail) CutFilter() {
	p.Filter = mongodb.FilterBsonM(p.Filter, []string{`_id`})
	delete(p.DataM, `_id`)
}

func (p *ConfigDetail) AddDelete() {
	p.Filter[`is_delete`] = 0
}

func (p *ConfigDetail) Convert2ObjectId() {
	if len(p.Filter[`_id`].(string)) > 0 {
		p.Filter[`_id`] = mongodb.ObjectIDFromHex(p.Filter[`_id`].(string))
	}
}

func (p *ConfigDetail) FindFields() {
	fieldList := bson.D{{Key: "_id", Value: 1},
		{Key: "project_id", Value: 1},
		{Key: "config_name", Value: 1},
		{Key: "codename", Value: 1},
		{Key: "content", Value: 1},
		{Key: "config_type", Value: 1},
		{Key: "theme_name", Value: 1},
		{Key: "use_callback", Value: 1},
		{Key: "create_time", Value: 1},
		{Key: "update_time", Value: 1}}
	p.FindOneOpts.SetProjection(fieldList)
}

func (p *ConfigDetail) GetListOrOne() {
	result, err := mongodb.FindOne("config", p.Filter, p.FindOneOpts)

	if err != nil {
		dhlog.Error(err.Error())
		p.Msg = utils.DebugMsg("config_config_detail_GetListOrOne FindOne 错误：" + p.Err.Error())
		p.MsgKey = "config_config_detail_GetListOrOne_failed"
		return
	}

	p.Result = result
}

func (p *ConfigDetail) TempPrepare() {
	method := "bson.M" // slice,pager
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

func (p *ConfigDetail) QueryOtherCollection() {

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
