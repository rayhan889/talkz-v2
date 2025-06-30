package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/http/middlewares"
	"github.com/rayhan889/talkz-v2/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Gin         *gin.Engine
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewApp(db *gorm.DB, client *redis.Client) *App {
	if config.Envs.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()

	r.Use(gin.Recovery())

	r.SetTrustedProxies(nil)

	r.Use(middlewares.SecurityHeaders())
	r.Use(middlewares.RateLimiter())
	r.Use(middlewares.CORS())

	return &App{
		Gin:         r,
		DB:          db,
		RedisClient: client,
	}
}

func (app *App) Run() {
	app.registerRoutes(app.Gin.Group("/api"))

	err := app.Gin.Run()
	if err != nil {
		panic(err)
	}
}

func (app *App) registerRoutes(api *gin.RouterGroup) {
	authService := InitializeAuthService(app.DB)
	// userService := InitializeUserService(app.DB)

	healthController := InitializeHealthController()
	authController := InitializeAuthController(authService)
	// userController := InitializeUserController(userService)

	v1 := api.Group("/v1")

	v1.GET("/health", healthController.HealthCheck)
	v1.POST("/auth/register", authController.Register)
	v1.POST("/auth/login", authController.Login)
	v1.GET("/auth/user", middlewares.Authenticate(authService), authController.User)
	v1.POST("/auth/refresh", authController.Refresh)
}
