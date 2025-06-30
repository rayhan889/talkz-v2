package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/helpers"
	"github.com/rayhan889/talkz-v2/pkg/logger"
)

func InternalServerError(c *gin.Context, err error) {
	logger.Log.Errorw("internal server error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	helpers.WriteJSONError(c, http.StatusInternalServerError, "Internal server error")
}

func ForbiddenError(c *gin.Context, err error) {
	logger.Log.Errorw("forbidden error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	helpers.WriteJSONError(c, http.StatusForbidden, "Forbidden")
}

func BadRequestError(c *gin.Context, err error) {
	logger.Log.Errorw("bad request error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	helpers.WriteJSONError(c, http.StatusBadRequest, "Bad request")
}

func ConflictError(c *gin.Context, err error) {
	logger.Log.Errorw("conflict error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	helpers.WriteJSONError(c, http.StatusConflict, "Conflict")
}

func NotFoundError(c *gin.Context, err error) {
	logger.Log.Errorw("not found error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	helpers.WriteJSONError(c, http.StatusNotFound, "Not found")
}

func UnauthorizedError(c *gin.Context, err error) {
	logger.Log.Errorw("unauthorized error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	helpers.WriteJSONError(c, http.StatusUnauthorized, "Unauthorized")
}

func UnauthorizedBasicError(c *gin.Context, err error) {
	logger.Log.Errorw("unauthorized basic error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	c.Header("WWW-Authenticate", "Basic realm=\"Restricted\"")

	helpers.WriteJSONError(c, http.StatusUnauthorized, "Unauthorized")
}

func RateLimitError(c *gin.Context, err error) {
	logger.Log.Errorw("rate limit error", "method", c.Request.Method, "path", c.Request.URL.Path, "error", err.Error())
	helpers.WriteJSONError(c, http.StatusTooManyRequests, "Rate limit exceeded")
}
