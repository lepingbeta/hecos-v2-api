package types


type CreateConfigParams struct {
	// 配置名
	ConfigName string `bson:"config_name" json:"config_name" validate:"required,min=2,max=50" `
	// 配置代号
	Codename string `bson:"codename" json:"codename" validate:"required,min=3,max=50" `
	// 修改后是否通知业务端（修改项目配置后调用回调函数）：true,false
	UseCallback bool `bson:"use_callback" json:"use_callback" validate:"boolean" `
	// 项目id
	ProjectId string `bson:"project_id" json:"project_id" validate:"required,mongoId,findInDb=needExists project _id create_config_project_id_find_in_db_err is_not_delete" `
	// 配置内容
	Content string `bson:"content" json:"content" validate:"required" `
	// 配置文件类型
	ConfigType string `bson:"config_type" json:"config_type" validate:"required" `
	// 主题名称
	ThemeName string `bson:"theme_name" json:"theme_name" validate:"required" `
}

