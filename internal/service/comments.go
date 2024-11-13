package service

import (
	"GoFeed/internal/repository"
)

type CommentService struct {
	repo repository.CommentRepo
}

func NewCommentService(cr repository.CommentRepo) *CommentService {
	return &CommentService{
		repo: cr,
	}
}
