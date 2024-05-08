package lib

import (
	"github.com/go-playground/validator/v10"
	dhvalidator "github.com/lepingbeta/go-common-v2-dh-validator"
)

func loadCommonValidator(validate *validator.Validate) {
	validate.RegisterValidation(`firstName`, dhvalidator.IsValidFirstName)
	validate.RegisterValidation(`chineseMobile`, dhvalidator.IsValidChineseMobile)
	validate.RegisterValidation(`lastName`, dhvalidator.IsValidLastName)
	validate.RegisterValidation(`nickname`, dhvalidator.IsValidNickname)
	validate.RegisterValidation(`gender`, dhvalidator.IsValidGender)
	validate.RegisterValidation(`image`, dhvalidator.IsValidImage)
	validate.RegisterValidation(`account`, dhvalidator.IsValidAccount)
	validate.RegisterValidation(`password`, dhvalidator.IsValidPassword)
}
