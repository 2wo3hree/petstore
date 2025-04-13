package service

import (
	"context"
	"repo/internal/model"
	"repo/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, u model.User) error
	GetByID(ctx context.Context, id string) (model.User, error)
	Update(ctx context.Context, u model.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]model.User, int, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, u model.User) error {
	return s.repo.Create(ctx, u)
}
func (s *userService) GetByID(ctx context.Context, id string) (model.User, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *userService) Update(ctx context.Context, u model.User) error {
	return s.repo.Update(ctx, u)
}
func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
func (s *userService) List(ctx context.Context, limit, offset int) ([]model.User, int, error) {
	return s.repo.List(ctx, limit, offset)
}
