package service

import (
	"context"
	"httpServer/repository"
)

type UserService interface {
	AddUser(ctx context.Context, u repository.User) (string, error)
	GetByID(ctx context.Context, userID string) (*repository.User, error)
	Update(ctx context.Context, user *repository.User) error
}

type userService struct {
	repo repository.UserRepo
}

func NewUserService(ur repository.UserRepo) *userService {
	return &userService{
		repo: ur,
	}
}

func (us *userService) AddUser(ctx context.Context, u repository.User) (string, error) {
	// Business logic

	id, err := us.repo.Add(ctx, u)
	if err != nil {
		return "", err
	}
	return id, nil
}
