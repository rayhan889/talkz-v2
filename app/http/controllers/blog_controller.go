package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/services"
)

type BlogController struct {
	blogService *services.BlogService
}

func NewBlogController(blogService *services.BlogService) *BlogController {
	return &BlogController{
		blogService: blogService,
	}
}

func (controller *BlogController) Compose(c *gin.Context) {

}
