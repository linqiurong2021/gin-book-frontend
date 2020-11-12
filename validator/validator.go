package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		v.RegisterValidation("bookabledate", bookableDate)
// 	}
