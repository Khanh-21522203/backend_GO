package repository

import (
	"context"
	"time"
)

type Post struct {
	ID               string    `json:"ID"`
	User_id          string    `json:"userID"`
	ContentText      string    `json:"contentText"`
	ContentImagePath string    `json:"contentImagePath"`
	CreatedAt        time.Time `json:"createdAt"`
}

type PostRepo interface {
	GetByID(ctx context.Context, postID string) (*Post, error)
	Add(ctx context.Context, p Post) error
	Update(ctx context.Context, p Post) error
	Delete(ctx context.Context, p Post) error
}
