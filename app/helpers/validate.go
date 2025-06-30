package helpers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	validatorPkg "github.com/rayhan889/talkz-v2/pkg/validator"
)

func ValidateRequest(c *gin.Context, request interface{}) {
	data, err := c.GetRawData()
	if err != nil {
		BadRequestError(c, err)
		return
	}

	err = json.Unmarshal(data, request)
	if err != nil {
		BadRequestError(c, err)
		return
	}

	err = validatorPkg.Validate.Struct(request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Validation failed",
		})
		return
	}
}

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
