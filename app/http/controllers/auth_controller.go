package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) Login(r *gin.Context) error {
	return nil
}

func (controller *AuthController) Logout(r *gin.Context) error {
	return nil
}

func (controller *AuthController) Register(r *gin.Context) error {
	return nil
}
