package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	FieldErrorResponse struct {
		Field   string      `json:"field"`
		Tag     string      `json:"tag"`
		Param   string      `json:"param,omitempty"`
		Value   interface{} `json:"value,omitempty"`
		Message string      `json:"message"`
	}
	ValidationError struct {
		FieldErrors []FieldErrorResponse `json:"field_errors"`
		OldFiedlds  interface{}          `json:"old_fields,omitempty"`
	}
)

func NewValidationError(c *gin.Context, fieldErrors []validator.FieldError, oldFields interface{}) {
	var errs []FieldErrorResponse
	for _, fe := range fieldErrors {
		errs = append(errs, FieldErrorResponse{
			Field:   fe.Field(),
			Tag:     fe.Tag(),
			Param:   fe.Param(),
			Value:   fe.Value(),
			Message: fe.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, &ValidationError{
		FieldErrors: errs,
		OldFiedlds:  oldFields,
	})
}
