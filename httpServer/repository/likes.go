package repository

import (
	"context"
	"time"
)

type Like struct {
	PostID    string    `json:"postID"`
	UserID    string    `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}

type LikeRepo interface {
	Add(ctx context.Context, userID string, postID string) error
	Count(ctx context.Context, postID string) (int, error)
}
