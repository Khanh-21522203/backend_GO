package repository

import (
	"context"
	"time"
)

type Comment struct {
	ID        string    `json:"ID"`
	PostID    string    `json:"postID"`
	UserID    string    `json:"userID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type CommentRepo interface {
	GetByPostID(ctx context.Context, pId string) ([]*Comment, error)
	Add(ctx context.Context, c Comment) error
	Update(ctx context.Context, c Comment) error
	Delete(ctx context.Context, c Comment) error
}
