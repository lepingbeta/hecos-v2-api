/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-09 18:26:50
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-11 00:15:07
 * @FilePath     : /hecos-v2-api/services/project/ProjectList/ProjectList.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package ProjectList

import (
	"fmt"
	"math"

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

type ProjectList struct {
	Params       t.ProjectListParams // 入参结构体版 （原始版）
	DataM        bson.M              // 入参bson.M版 (入库用)
	SliceOfDataM []bson.M            // 入参slice版
	Filter       bson.M              // 入参bson.M版 (查询用)
	DataD        bson.D              // 入参bson.D版 (入库用)
	C            *gin.Context
	Result       any
	Msg          string
	MsgKey       string
	Err          error
	DocID        primitive.ObjectID
}

func (p *ProjectList) ProjectList() {

	p.GetListWithPager()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}

	fmt.Println("Hello, my name is")
}

func (p *ProjectList) GetListWithPager() {
	if true {
		p.Filter["is_delete"] = 0
	}
	// 从请求中获取了页码和每页大小
	page := int64(p.Filter["page"].(int32))
	pageSize := int64(p.Filter["page_size"].(int32))
	delete(p.Filter, "page")
	delete(p.Filter, "page_size")

	// 计算跳过的文档数
	skip := (page - 1) * pageSize

	opts := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(skip))
	fieldList := bson.D{{Key: "_id", Value: 1},
		{Key: "project_name", Value: 1},
		{Key: "update_callback", Value: 1},
		{Key: "accessId", Value: 1},
		{Key: "create_time", Value: 1}}
	opts.SetProjection(fieldList)
	result, err := mongodb.FindList("project", p.Filter, opts)

	if err != nil {
		dhlog.Error(err.Error())
		p.Msg = utils.DebugMsg("读取数据失败：" + p.Err.Error())
		p.MsgKey = "project_project_list_GetListWithPager_failed"
		return
	}

	count, err := mongodb.Count("project", p.Filter)
	if err != nil {
		p.Err = err
		p.Msg = utils.DebugMsg("project_project_list_GetListWithPager count 错误" + err.Error())
		p.MsgKey = "project_project_list_GetListWithPager_query_count_error"
		dhlog.Error(p.Msg)
		return
	}

	var finalResult httpTypes.DataList
	finalResult.Page = page
	finalResult.Total = int64(math.Ceil(float64(count) / float64(pageSize)))
	finalResult.List = result

	p.Result = finalResult
}

// {{占位符 composition}}
