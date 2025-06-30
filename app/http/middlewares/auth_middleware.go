package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/constants"
	"github.com/rayhan889/talkz-v2/app/services"
	"github.com/rayhan889/talkz-v2/pkg/logger"
)

func Authenticate(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, constants.AccessTokenPrefix) {
			logger.Log.Error(constants.MissingAccessTokenError)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": constants.MissingAccessTokenError,
			})
			return
		}

		token := strings.TrimSpace(authHeader[len(constants.AccessTokenPrefix):])

		user, err := authService.ValidateAccessToken(token)

		if err != nil {
			logger.Log.Error(constants.InvalidAccessToken)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": constants.InvalidAccessToken,
			})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
