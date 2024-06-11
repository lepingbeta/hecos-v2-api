/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-10 14:35:40
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-11 00:41:59
 * @FilePath     : /hecos-v2-api/types/DeleteProjectParams.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package types

type DeleteProjectParams struct {
	// 项目id
	ProjectId string `bson:"_id" json:"_id" validate:"required,mongoId,findInDb=needExists project _id delete_project__id_find_in_db_err is_not_delete" `
}
