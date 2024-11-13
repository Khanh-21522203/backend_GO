package service

import (
	"GoFeed/internal/repository"
)

type PostService struct {
	repo repository.PostRepo
}

func NewPostService(pr repository.PostRepo) *PostService {
	return &PostService{
		repo: pr,
	}
}
