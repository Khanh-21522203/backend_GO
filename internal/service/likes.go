package service

import (
	"GoFeed/internal/repository"
)

type LikeService struct {
	repo repository.LikeRepo
}

func NewLikeRepo(pr repository.LikeRepo) *LikeService {
	return &LikeService{
		repo: pr,
	}
}
