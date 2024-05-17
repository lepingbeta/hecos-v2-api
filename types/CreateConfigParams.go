package types


type CreateConfigParams struct {
	// 配置名
	ConfigName string `bson:"config_name" json:"config_name" validate:"required,min=3,max=50" `
	// 项目id
	ProjectId string `bson:"project_id" json:"project_id" validate:"required,mongoId" `
	// 配置内容
	Content string `bson:"content" json:"content" validate:"required" `
	// 配置文件类型
	ConfigType string `bson:"config_type" json:"config_type" validate:"required" `
	// 主题名称
	ThemeName string `bson:"theme_name" json:"theme_name" validate:"required" `
}

