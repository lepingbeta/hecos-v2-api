package types


type ConfigDetailParams struct {
	// 配置id
	Id string `bson:"_id" json:"_id" form:"_id" validate:"required,mongoId,findInDb=needExists config _id config_detail__id_find_in_db_err is_not_delete" `
}

