package health

import (
	"github.com/gin-gonic/gin"
)

func (h *HealthHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/health", h.healthCheck)
}
