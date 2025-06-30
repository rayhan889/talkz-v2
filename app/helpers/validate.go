package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(request interface{}) []validator.FieldError {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(request)

	if err != nil {
		var errors []validator.FieldError

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err)
		}

		return errors
	}

	return nil
}
