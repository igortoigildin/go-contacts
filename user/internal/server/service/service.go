package service

import (
	"context"

	"github.com/igortoigildin/go-contacts/user/internal/server/models"
)

type UserService struct{}

func New(ctx context.Context) *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) error {
	// TODO: implement
	return nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (models.User, error) {
	// TODO: implement
	return models.User{}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string) (models.User, error) {
	// TODO: implement
	return models.User{}, nil
}

func (s *UserService) SearchUsers(ctx context.Context, query string) ([]*models.User, error) {
	// TODO: implement
	return []*models.User{}, nil
}
