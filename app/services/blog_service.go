package services

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/rayhan889/talkz-v2/app/models"
	"github.com/rayhan889/talkz-v2/app/repositories"
	"github.com/rayhan889/talkz-v2/pkg/logger"
)

type BlogService struct {
	blogRepostiory *repositories.BlogRepository
}

func NewBlogService(blogRepostiory *repositories.BlogRepository) *BlogService {
	return &BlogService{
		blogRepostiory: blogRepostiory,
	}
}

func (service *BlogService) CreateBlog(
	title string,
	content string,
	authorId uuid.UUID,
) (*models.Blog, error) {
	slug, err := service.GenerateSlug(title)

	if err != nil {
		return nil, err
	}

	blog := models.Blog{
		ID:       uuid.New(),
		Title:    title,
		Slug:     slug,
		Content:  content,
		AuthorID: authorId,
	}

	err = service.blogRepostiory.Create(&blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (service *BlogService) GetCountBySlug(slug string) (int64, error) {
	blogs, err := service.blogRepostiory.FindBySlug(slug)
	logger.Log.Infof("Blogs: %s", blogs)

	if err != nil {
		return 0, err
	}

	return int64(len(blogs)), nil
}

func (service *BlogService) GenerateSlug(title string) (string, error) {
	slug := strings.ToLower(title)
	slug = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == ' ' {
			return r
		}
		return -1
	}, slug)

	slug = strings.ReplaceAll(slug, " ", "-")

	blogs, err := service.GetCountBySlug(slug)

	if err != nil {
		return "", err
	}

	if blogs > 0 {
		return fmt.Sprintf("%s-%d", slug, blogs), nil
	}

	return slug, nil
}
