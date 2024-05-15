package types


type CreateConfigParams struct {
	// 配置名
	ConfigName string `bson:"config_name" json:"config_name" validate:"required,min=3,max=50" `
	// 项目id
	ProjectId string `bson:"project_id" json:"project_id" validate:"required,mongoId" `
	// 配置内容
	Content string `bson:"content" json:"content" validate:"required" `
}

