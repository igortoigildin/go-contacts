package user

import (
	"context"

	"github.com/igortoigildin/go-contacts/user/internal/server/service"
	desc "github.com/igortoigildin/go-contacts/user/pkg/api/users"
)

type Implementation struct {
	userService *service.UserService
	desc.UnimplementedUserServiceServer
}

func NewImplementation(userService *service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

func (i *Implementation) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.User, error) {
	err := i.userService.CreateUser(ctx, req.GetName(), req.GetEmail())
	if err != nil {
		return nil, err
	}

	return &desc.User{}, nil
}

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.User, error) {
	user, err := i.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.User{
		Id:    user.ID,
		Name:  user.Nickname,
		Email: user.Email,
	}, nil
}

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.User, error) {
	user, err := i.userService.UpdateUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.User{
		Id:    user.ID,
		Name:  user.Nickname,
		Email: user.Email,
	}, nil
}

func (i *Implementation) SearchUsers(ctx context.Context, req *desc.SearchUsersRequest) (*desc.SearchUsersResponse, error) {
	_, err := i.userService.SearchUsers(ctx, req.GetQuery())
	if err != nil {
		return nil, err
	}
	// TODO: implement
	return &desc.SearchUsersResponse{
		Users: []*desc.User{},
	}, nil
}
