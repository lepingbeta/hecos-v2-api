package types


type DeleteConfigParams struct {
	// 配置id
	Id string `bson:"_id" json:"_id" validate:"required,mongoId,findInDb=needExists config _id delete_config__id_find_in_db_err is_not_delete" `
}

