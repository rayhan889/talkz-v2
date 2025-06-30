package services

import "github.com/rayhan889/talkz-v2/app/repositories"

type BlogService struct {
	blogRepostiory *repositories.BlogRepository
}

func NewBlogService(blogRepostiory *repositories.BlogRepository) *BlogService {
	return &BlogService{
		blogRepostiory: blogRepostiory,
	}
}
