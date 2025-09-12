// api/validation/mobile.go
package validation

import (
    "github.com/go-playground/validator/v10"
    "user-crud/common"
)

func RegisterMobileValidation(v *validator.Validate) {
    v.RegisterValidation("mobile", IranianMobileNumberValidator)
}

func IranianMobileNumberValidator(fl validator.FieldLevel) bool {
    value, ok := fl.Field().Interface().(string)
    if !ok {
        return false
    }
    return common.IranianMobileNumberValidate(value)
}
