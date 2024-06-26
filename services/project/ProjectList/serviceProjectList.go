package ProjectList

import (
	httpTypes "github.com/lepingbeta/go-common-v2-dh-http/types"
	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	mongodb "github.com/lepingbeta/go-common-v2-dh-mongo"
	utils "github.com/lepingbeta/go-common-v2-dh-utils"
	"go.mongodb.org/mongo-driver/bson"
	"math"
	// {{占位符 import}}
)

func (p *ProjectList) ProjectList() {

	p.AddDelete()
	if p.Err != nil {
		return
	}
	p.FindFields()
	if p.Err != nil {
		return
	}
	p.FindSort()
	if p.Err != nil {
		return
	}
	p.GetListWithPager()
	if p.Err != nil {
		return
	} // {{占位符 composition caller}}
}

func (p *ProjectList) AddDelete() {
	p.Filter[`is_delete`] = 0
}

func (p *ProjectList) FindFields() {
	fieldList := bson.D{{Key: "_id", Value: 1},
		{Key: "project_name", Value: 1},
		{Key: "update_callback", Value: 1},
		{Key: "accessId", Value: 1},
		{Key: "create_time", Value: 1}}
	p.FindOpts.SetProjection(fieldList)
}

func (p *ProjectList) FindSort() {
	if mongodb.HasKey(p.Filter, "sort_fields") {
		sortClause, err := mongodb.ParseSortString(p.Filter["sort_fields"].(string))
		if err != nil {
			p.Err = err
			p.Msg = utils.DebugMsg("project_project_list_FindSort error" + err.Error())
			p.MsgKey = "project_project_list_FindSort_error"
			dhlog.Error(p.Msg)
			return
		}
		delete(p.Filter, "sort_fields")
		p.FindOpts.SetSort(sortClause)
	}
}

// 分页逻辑
func (p *ProjectList) pagination() (int64, int64) {
	// 从请求中获取了页码和每页大小
	page := int64(p.Filter["page"].(int32))
	pageSize := int64(p.Filter["page_size"].(int32))
	delete(p.Filter, "page")
	delete(p.Filter, "page_size")

	// 计算跳过的文档数
	skip := (page - 1) * pageSize

	p.FindOpts.SetLimit(int64(pageSize)).SetSkip(int64(skip))
	return page, pageSize
}

// 查询主逻辑
func (p *ProjectList) GetListWithPager() {
	page, pageSize := p.pagination()

	result, err := mongodb.FindList("project", p.Filter, p.FindOpts)

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

	// 组合分页信息
	var finalResult httpTypes.DataList
	finalResult.Page = page
	finalResult.Total = int64(math.Ceil(float64(count) / float64(pageSize)))
	finalResult.List = result

	p.Result = finalResult
}

// {{占位符 composition}}
