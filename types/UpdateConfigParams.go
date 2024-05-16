package types


type UpdateConfigParams struct {
	// 配置名
	ConfigName string `bson:"config_name" json:"config_name" validate:"required,min=3,max=50" `
	// 配置内容
	Content string `bson:"content" json:"content" validate:"required" `
	// 配置id
	Id string `bson:"_id" json:"_id" validate:"required,mongoId,findInDb=needExists config _id update_config__id_find_in_db_err zero" `
}

