package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/http/middlewares"
	"github.com/rayhan889/talkz-v2/config"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type App struct {
	Gin         *gin.Engine
	DB          *gorm.DB
	RedisClient *redis.Client
	Dialer      *gomail.Dialer
}

func NewApp(db *gorm.DB, client *redis.Client, dialer *gomail.Dialer) *App {
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
		Dialer:      dialer,
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
	authService := InitializeAuthService(app.DB, app.Dialer)
	// userService := InitializeUserService(app.DB)
	blogService := InitializeBlogService(app.DB, app.RedisClient)

	healthController := InitializeHealthController()
	authController := InitializeAuthController(authService)
	// userController := InitializeUserController(userService)
	blogController := InitializeBlogController(blogService)

	v1 := api.Group("/v1")

	v1.GET("/health", healthController.HealthCheck)

	v1.POST("/auth/register", authController.Register)
	v1.POST("/auth/login", authController.Login)
	v1.GET("/auth/user", middlewares.Authenticate(authService), authController.User)
	v1.POST("/auth/refresh", authController.Refresh)

	v1.GET("/blogs/feeds", middlewares.Authenticate(authService), blogController.Feeds)
	v1.POST("/blogs/compose", middlewares.Authenticate(authService), blogController.Compose)
}
