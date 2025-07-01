package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/rayhan889/talkz-v2/app/models"
	"github.com/rayhan889/talkz-v2/app/repositories"
	"github.com/rayhan889/talkz-v2/config"
	"github.com/rayhan889/talkz-v2/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type BlogService struct {
	blogRepostiory *repositories.BlogRepository
	redis          *redis.Client
}

func NewBlogService(blogRepostiory *repositories.BlogRepository, redis *redis.Client) *BlogService {
	return &BlogService{
		blogRepostiory: blogRepostiory,
		redis:          redis,
	}
}

func (service *BlogService) GetFeeds() ([]models.Blog, error) {
	blogs, err := service.blogRepostiory.FindAll()

	if err != nil {
		return nil, err
	}

	cachedBlogs, err := service.GetCachedBlogs("blogs:feed")
	if err == nil && len(cachedBlogs) > 0 {
		return cachedBlogs, nil
	}

	service.SetCacheBlogs(blogs, "blogs:feed")

	return blogs, nil
}

func (service *BlogService) GetCachedBlogs(key string) ([]models.Blog, error) {
	val, err := service.redis.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	var blogs []models.Blog
	if err := json.Unmarshal([]byte(val), &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (service *BlogService) SetCacheBlogs(blogs []models.Blog, key string) error {
	data, err := json.Marshal(blogs)

	if err != nil {
		return err
	}

	ttl := time.Duration(config.Envs.Redis.Duration) * time.Minute
	service.redis.Set(context.Background(), key, data, ttl)

	return nil
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

	go func() {
		if err := service.redis.Del(context.Background(), "blogs:feed").Err(); err != nil {
			logger.Log.Errorf("Failed to invalidate blogs:feed cache: %v", err)
		}
	}()

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
