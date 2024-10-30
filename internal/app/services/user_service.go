package services

import (
	"context"

	"github.com/northwindman/book-shop/internal/app/domain"
)

// UserService is a user service
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s UserService) GetUser(ctx context.Context, username string) (domain.User, error) {
	return s.repo.GetUser(ctx, username)
}

func (s UserService) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
