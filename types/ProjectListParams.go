package types


type ProjectListParams struct {
	// 当前页
	Page int `bson:"page" json:"page" form:"page" validate:"required,gt=0" `
	// 每页显示条目数
	PageSize int `bson:"page_size" json:"page_size" form:"page_size" validate:"required,gt=0" `
}

