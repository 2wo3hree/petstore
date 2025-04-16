package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user models.User) error
	GetByUsername(ctx context.Context, username string) (models.User, error)
	Delete(ctx context.Context, username string) error
	Update(ctx context.Context, username string, updated models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(ctx, user)
}

func (s *userService) GetByUsername(ctx context.Context, username string) (models.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

func (s *userService) Delete(ctx context.Context, username string) error {
	return s.repo.Delete(ctx, username)
}

func (s *userService) Update(ctx context.Context, username string, updated models.User) error {
	return s.repo.Update(ctx, username, updated)
}
