// src/api/validation/init.go
package validation

import "github.com/go-playground/validator/v10"

func RegisterCustomValidators(v *validator.Validate) {
	// ثبت validator شماره موبایل
	v.RegisterValidation("mobile", IranianMobileNumberValidator)
	
	// ثبت validator رمز عبور
	v.RegisterValidation("password", PasswordValidator)
	
	// ثبت validatorهای دیگر...
}