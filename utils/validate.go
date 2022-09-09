package utils

import (
	"github/koybigino/getting-started-fiber/models"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
	Message     string
}

var validate = validator.New()

func ValidateStruct(user models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Message = "Invalid email, please enter a right one like this  -> 'ali@ggd.dsf' !"
			errors = append(errors, &element)
		}
	}
	return errors
}
