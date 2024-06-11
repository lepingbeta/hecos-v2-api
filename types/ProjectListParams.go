/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-09 06:17:44
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-12 06:45:13
 * @FilePath     : /hecos-v2-api/types/ProjectListParams.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package types

type ProjectListParams struct {
	// 当前页
	Page int `bson:"page" json:"page" form:"page" validate:"required,gt=0" `
	// 每页显示条目数
	PageSize int `bson:"page_size" json:"page_size" form:"page_size" validate:"required,gt=0" `
	// 要排序的字段，可以为空
	SortFields string `bson:"sort_fields" json:"sort_fields" form:"sort_fields" validate:"omitempty" `
}
