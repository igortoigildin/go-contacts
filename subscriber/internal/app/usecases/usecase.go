package usecases

import (
	"context"
)

type Usecase interface {
	MakeFriendRequest(ctx context.Context, req *FriendRequestDTO) error
	AcceptFriendRequest(ctx context.Context, req *FriendAcceptDTO) error
	RejectFriendRequest(ctx context.Context, req *FriendRejectDTO) error
	RemoveFriend(ctx context.Context, req *RemoveFriendDTO) error
	ListFriends(ctx context.Context, req *ListFriendsDTO) ([]string, error)
}

var (
	_ Usecase = (*usecase)(nil)
)

type usecase struct {
	Deps
}

func NewUsecase(deps Deps) *usecase {
	return &usecase{
		Deps: deps,
	}
}

type Deps struct{}
