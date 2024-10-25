package service

import (
	"httpServer/repository"
)

type PostService struct {
	repo repository.PostRepo
}

func NewPostService(pr repository.PostRepo) *PostService {
	return &PostService{
		repo: pr,
	}
}
