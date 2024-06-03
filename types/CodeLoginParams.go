package types


type CodeLoginParams struct {
	// iuc的登录code
	Code string `bson:"code" json:"code" form:"code" validate:"required" `
}

