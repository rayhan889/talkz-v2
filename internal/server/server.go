package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/internal/handler/health"
	"github.com/rayhan889/talkz-v2/internal/middleware"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Start(env string) error {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()

	r.Use(gin.Recovery())

	r.SetTrustedProxies(nil)

	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.RateLimiter())

	v1 := r.Group("/api/v1")

	s.bindRoutes(v1)

	return r.Run(s.addr)
}

func (s *APIServer) bindRoutes(r *gin.RouterGroup) {
	healthHanlder := health.NewHealthHandler()
	healthHanlder.RegisterRoutes(r)
}
