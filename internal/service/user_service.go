package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, u models.User) (int64, error)
	List(ctx context.Context, limit, offset int) ([]models.User, int, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	GetWithRentals(ctx context.Context, id int) (models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, u models.User) (int64, error) {
	return s.repo.Create(ctx, u)
}

func (s *userService) List(ctx context.Context, limit, offset int) ([]models.User, int, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *userService) GetByID(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) GetWithRentals(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetWithRentals(ctx, id)
}
