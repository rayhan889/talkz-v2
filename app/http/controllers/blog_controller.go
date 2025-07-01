package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/constants"
	"github.com/rayhan889/talkz-v2/app/exceptions"
	"github.com/rayhan889/talkz-v2/app/helpers"
	"github.com/rayhan889/talkz-v2/app/http/requests"
	"github.com/rayhan889/talkz-v2/app/http/responses"
	"github.com/rayhan889/talkz-v2/app/models"
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
	user := c.MustGet("user").(*models.User)

	composeRequest := new(requests.ComposeBlogRequest)

	err := c.BindJSON(composeRequest)

	if err != nil {
		exceptions.BadRequestError(c, errors.New(constants.ErrorInvalidRequestBody))
		return
	}

	errs := helpers.ValidateStruct(composeRequest)
	if errs != nil {
		exceptions.NewValidationError(c, errs, composeRequest)
		return
	}

	blog, err := controller.blogService.CreateBlog(composeRequest.Title, composeRequest.Content, user.ID)

	if err != nil {
		exceptions.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"data": responses.ComposeBlogRespone{
			ID:      blog.ID.String(),
			Title:   blog.Title,
			Content: blog.Content,
		},
		"errors": nil,
	})
}
