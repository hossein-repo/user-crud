package validation

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IranianMobileNumberValidator(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()

	// الگوی شماره موبایل ایرانی
	pattern := `^09[0-9]{9}$`

	matched, _ := regexp.MatchString(pattern, mobile)

	return matched
}

func RegisterMobileValidation(v *validator.Validate) {
	fmt.Println("🚀RegisterMobileValidation به به  residim rdgister .")

	v.RegisterValidation("mobile", IranianMobileNumberValidator)
}
