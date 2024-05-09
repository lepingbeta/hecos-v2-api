package CreateProject

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	t "tangxiaoer.shop/dahe/hecos-v2-api/types"
)


func CreateProject(user t.CreateProjectParams) (map[string]interface{}, string, string, error) {

	user, msg, msgKey, err := CreateProjectPre(user)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	result, err := mongodb.InsertOneWithCreateTime("project", user)

	if err != nil {
		dhlog.Error(err.Error())
		return nil, "数据入库失败", "project_create_project_insert_to_db_failed", err
	}

	// 获取并打印 _id
	// 获取并尝试将 _id 转换为 primitive.ObjectID
	docID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		dhlog.Error("Expected the inserted document ID to be a primitive.ObjectID")
		return nil, "Expected the inserted document ID to be a primitive.ObjectID", "project_create_project_insert_id_error", err
	}

	finalResult, msg, msgKey, err := CreateProjectPost(user, docID)
	if err != nil {
		dhlog.Error(err.Error())
		return nil, msg, msgKey, err
	}

	return finalResult, msg, "", nil
}
