package types


type CreateProjectParams struct {
	// 邮箱
	ProjectName string `bson:"project_name" json:"project_name" validate:"required,max=50" `
}

