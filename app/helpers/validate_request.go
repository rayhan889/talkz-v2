package helpers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/pkg/validator"
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

	err = validator.Validate.Struct(request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Validation failed",
		})
		return
	}
}
