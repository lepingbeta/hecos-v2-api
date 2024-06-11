/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-06-11 01:25:06
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-11 01:25:17
 * @FilePath     : /hecos-v2-api/types/SearchProjectListParams.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package types

type SearchProjectListParams struct {
	// 项目id
	ProjectId string `bson:"_id" json:"_id" form:"_id" validate:"omitempty,mongoId" `
}
