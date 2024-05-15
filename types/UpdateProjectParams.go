package types


type UpdateProjectParams struct {
	// 项目名
	ProjectName string `bson:"project_name" json:"project_name" validate:"required,min=3,max=50" `
	// 项目id
	ProjectId string `bson:"project_id" json:"project_id" validate:"required,mongoId" `
}

