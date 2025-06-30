package middlewares

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/config"
)

func CORS() gin.HandlerFunc {
	allowedHeaders := strings.Split(config.Envs.Cors.AllowHeaders, ",")
	allowedMethods := strings.Split(config.Envs.Cors.AllowMethods, ",")
	exposeHeaders := strings.Split(config.Envs.Cors.ContentLength, ",")
	maxAgeTime := time.Duration(config.Envs.Cors.MaxAge) * time.Hour

	return cors.New(cors.Config{
		AllowMethods:     allowedMethods,
		AllowHeaders:     allowedHeaders,
		ExposeHeaders:    exposeHeaders,
		AllowCredentials: config.Envs.Cors.AllowCredentials,
		AllowOriginFunc: func(origin string) bool {
			return strings.Contains(config.Envs.Cors.AllowOrigins, origin)
		},
		MaxAge: maxAgeTime,
	})
}
