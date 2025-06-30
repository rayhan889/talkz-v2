package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/constants"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(c *gin.Context, status int, data any) {
	c.Header("Content-Type", constants.JSONContentType)
	c.JSON(status, data)
}

func ReadJSON(c *gin.Context, data any) error {
	return c.ShouldBindJSON(data)
}

func WriteJSONError(c *gin.Context, status int, message string) {
	type errResponse struct {
		Error string `json:"error"`
	}
	c.JSON(status, &errResponse{
		Error: message,
	})
}
