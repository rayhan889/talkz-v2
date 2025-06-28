package health

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v5"
	healthPostgres "github.com/hellofresh/health-go/v5/checks/postgres"
	healthRedis "github.com/hellofresh/health-go/v5/checks/redis"
	"github.com/rayhan889/talkz-v2/config"
)

func (h *HealthHandler) healthCheck(c *gin.Context) {

	version := config.Envs.App.Version

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	instance, _ := health.New(health.WithComponent(health.Component{
		Name:    "talkz",
		Version: version,
	}), health.WithSystemInfo())

	instance.Register(health.Config{
		Name:      "postgres",
		Timeout:   time.Second * 2,
		SkipOnErr: false,
		Check: healthPostgres.New(healthPostgres.Config{
			DSN: config.Envs.DB.Address,
		}),
	})

	instance.Register(health.Config{
		Name:      "redis",
		Timeout:   time.Second * 2,
		SkipOnErr: true,
		Check: healthRedis.New(healthRedis.Config{
			DSN: config.Envs.Redis.Address,
		}),
	})

	result := instance.Measure(c.Request.Context())

	c.JSON(200, result)
}
