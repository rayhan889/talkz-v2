package middlewares

import (
	"strings"
	"time"

	"slices"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/config"
)

func CORS() gin.HandlerFunc {
	allowedHeaders := strings.Split(config.Cors.AllowHeaders, ",")
	allowedMethods := strings.Split(config.Cors.AllowMethods, ",")
	exposeHeaders := strings.Split(config.Cors.ContentLength, ",")
	maxAgeTime := time.Duration(config.Cors.MaxAge) * time.Hour

	allowedOrigins := strings.Split(config.Cors.AllowOrigins, ",")
	for i := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(allowedOrigins[i])
	}

	return cors.New(cors.Config{
		AllowMethods:     allowedMethods,
		AllowHeaders:     allowedHeaders,
		ExposeHeaders:    exposeHeaders,
		AllowCredentials: config.Cors.AllowCredentials,
		AllowOriginFunc: func(origin string) bool {
			return slices.Contains(allowedOrigins, origin)
		},
		MaxAge: maxAgeTime,
	})
}
