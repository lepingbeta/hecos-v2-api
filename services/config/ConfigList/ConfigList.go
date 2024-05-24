package ConfigList

import (
	"github.com/gin-gonic/gin"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)

func ConfigList(params t.ConfigListParams, c *gin.Context) ([]bson.M, string, string, error) {

	filter, msg, msgKey, err := ConfigListPre(params, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.FindList("config", filter)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, err.Error(), "config_config_list_find_list_error", err
	}

	// 后置处理器
	filter, result, msg, msgKey, err = postProcessing(params, filter, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	finalResult, msg, msgKey, err := ConfigListPost(params, filter, result, c)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, msgKey, err
}

func postProcessing(params t.ConfigListParams, filter bson.M, result []bson.M, c *gin.Context) (bson.M, []bson.M, string, string, error) {
	filter, result, msg, msgKey, err := findWithOneByOne(params, filter, result, c)
	if err != nil {
		return filter, result, msg, msgKey, nil
	}
	// {{占位符 postProcessing}}
	return filter, result, "Post Processing Success", "config_config_list_post_processing_success", nil
}

func findWithOneByOne(params t.ConfigListParams, filter bson.M, result []bson.M, c *gin.Context) (bson.M, []bson.M, string, string, error) {
	var sourceValueList []primitive.ObjectID
	// 遍历切片
	for _, item := range result {
		// 使用item["project_id"]来访问每个映射中的project_id字段
		sourceValue, err := primitive.ObjectIDFromHex(item["project_id"].(string))
		if err != nil {
			return filter, result, err.Error(), "config_config_list_findWithOneByOne_source_value_error", nil
		}

		sourceValueList = append(sourceValueList, sourceValue)
	}
	filter2 := bson.M{"_id": bson.M{"$in": sourceValueList}}
	fieldList := bson.D{{"project_name", 1}, {"update_callback", 1}, {"_id", 1}}
	// 创建Find选项，设置Projection
	findOptions := options.Find()
	findOptions.SetProjection(fieldList)
	result2, err := mongodb.FindList("project", filter2, findOptions)
	if err != nil {
		return filter, result, err.Error(), "config_config_list_findWithOneByOne_source_find_list_error", nil
	}

	// 创建一个映射，以project_id为键，以project_name为值
	idToDetailsMap := make(map[primitive.ObjectID]bson.M)
	for _, res2 := range result2 {
		keyId, _ := primitive.ObjectIDFromHex(res2["_id"].(string))

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

	for i, item := range result {
		keyId, _ := primitive.ObjectIDFromHex(item["project_id"].(string))
		if details, exists := idToDetailsMap[keyId]; exists {
			// 将details中的字段添加到result[i]中
			for key, value := range details {
				// 检查result[i]中是否存在该字段，如果存在则添加
				if _, exists := result[i][key]; !exists {
					result[i][key] = value
				}
			}
		}
	}

	return filter, result, "findWithOneByOne查询成功", "config_config_list_findWithOneByOne_success", nil
}

// {{占位符 processer}}
