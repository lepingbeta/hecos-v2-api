package types


type SendInviteCodeParams struct {
	// 邮箱
	Email string `bson:"email" json:"email" validate:"required,email" `
}

