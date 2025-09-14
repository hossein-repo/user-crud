// src/api/validation/mobile.go
package validation

import (
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// IranianMobileNumberValidator اعتبارسنجی شماره موبایل ایرانی
func IranianMobileNumberValidator(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	// الگوی شماره موبایل ایرانی
	pattern := `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

// PasswordValidator اعتبارسنجی رمز عبور
func PasswordValidator(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	// حداقل ۸ کاراکتر
	if len(value) < 8 {
		return false
	}

	// دارای حروف بزرگ و کوچک
	var hasUpper, hasLower, hasDigit bool
	for _, char := range value {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}
