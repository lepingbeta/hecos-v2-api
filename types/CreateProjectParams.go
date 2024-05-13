package types


type CreateProjectParams struct {
	// 邮箱
	ProjectName string `bson:"project_name" json:"project_name" validate:"required,max=50,findInDb=needNotExists project project_name create_project_project_name_find_in_db_err zero" `
}

