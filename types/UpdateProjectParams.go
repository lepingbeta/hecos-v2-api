package types

type UpdateProjectParams struct {
	// 项目名
	ProjectName string `bson:"project_name" json:"project_name" validate:"required,min=3,max=50" `
	// 项目id
	ProjectId string `bson:"_id" json:"_id" validate:"required,mongoId,findInDb=needExists project _id update_project__id_find_in_db_err is_not_delete" `
	// 修改配置后回调接口，如果有的话
	UpdateCallback string `bson:"update_callback" json:"update_callback" validate:"" `
}
