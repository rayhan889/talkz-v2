//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/rayhan889/talkz-v2/app/http/controllers"
	"github.com/rayhan889/talkz-v2/app/repositories"
	"github.com/rayhan889/talkz-v2/app/services"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitializeApp(db *gorm.DB, client *redis.Client) *App {
	panic(wire.Build(
		NewApp,
	))
}

func InitializeUserRepository(db *gorm.DB) *repositories.UserRepository {
	panic(wire.Build(
		repositories.NewUserRepository,
	))
}

func InitializeBlogRepository(db *gorm.DB) *repositories.BlogRepository {
	panic(wire.Build(
		repositories.NewBlogRepository,
	))
}

func InitializeUserService(db *gorm.DB) *services.UserService {
	panic(wire.Build(
		services.NewUserService,
		InitializeUserRepository,
	))
}

func InitializeBlogService(db *gorm.DB) *services.BlogService {
	panic(wire.Build(
		services.NewBlogService,
		InitializeBlogRepository,
	))
}

func InitializeAuthService(db *gorm.DB) *services.AuthService {
	panic(wire.Build(
		services.NewAuthService,
		InitializeUserService,
	))
}

func InitializeUserController(userService *services.UserService) *controllers.UserController {
	panic(wire.Build(
		controllers.NewUserController,
	))
}

func InitializeBlogController(blogService *services.BlogService) *controllers.BlogController {
	panic(wire.Build(
		controllers.NewBlogController,
	))
}

func InitializeAuthController(authService *services.AuthService) *controllers.AuthController {
	panic(wire.Build(
		controllers.NewAuthController,
	))
}

func InitializeHealthController() *controllers.HealthController {
	panic(wire.Build(
		controllers.NewHealthController,
	))
}
