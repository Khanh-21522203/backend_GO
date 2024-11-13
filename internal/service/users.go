package service

import (
	"GoFeed/internal/repository"
	"context"
)

type UserService interface {
	AddUser(ctx context.Context, u *repository.User) (string, error)
	GetByID(ctx context.Context, userID string) (*repository.User, error)
	Update(ctx context.Context, u *repository.User) (string, error)
}

type userService struct {
	UserService
	repo repository.UserRepo
}

func NewUserService(ur repository.UserRepo) *userService {
	return &userService{
		repo: ur,
	}
}

func (us *userService) AddUser(ctx context.Context, u *repository.User) (string, error) {
	// Business logic

	id, err := us.repo.Add(ctx, u)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (us *userService) GetByID(ctx context.Context, userID string) (*repository.User, error) {
	user, err := us.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Update(ctx context.Context, u *repository.User) (string, error) {
	userID, err := us.repo.Update(ctx, u)
	if err != nil {
		return "", err
	}
	return userID, nil
}
