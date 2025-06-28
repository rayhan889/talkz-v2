package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) Index(r *gin.Context) error {
	return nil
}
