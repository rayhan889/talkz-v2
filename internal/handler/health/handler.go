package health

import (
	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v5"
	"github.com/rayhan889/talkz-v2/config"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) healthCheck(c *gin.Context) {

	version := config.Envs.App.Version

	instance, _ := health.New(health.WithComponent(health.Component{
		Name:    "talkz",
		Version: version,
	}))

	result := instance.Measure(c.Request.Context())

	c.JSON(200, result)
}
