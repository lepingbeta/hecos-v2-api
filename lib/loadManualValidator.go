/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-06 05:06:20
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-08 20:48:53
 * @FilePath     : /hecos-v2-api/lib/loadManualValidator.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package lib

import (
	"github.com/go-playground/validator/v10"
	dh_validator_manual "github.com/lepingbeta/go-common-v2-dh-validator-manual"
)

func loadManualValidator(validate *validator.Validate) {
	validate.RegisterValidation(`findInDb`, dh_validator_manual.IsValidfindInDb)
}
