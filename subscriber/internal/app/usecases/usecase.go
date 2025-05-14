package usecases

import (
	"context"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"

	subscriber_models "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/models"
)

type Usecase interface {
	MakeFriendRequest(ctx context.Context, req *FriendRequestDTO) error
	AcceptFriendRequest(ctx context.Context, req *FriendAcceptDTO) error
	RejectFriendRequest(ctx context.Context, req *FriendRejectDTO) error
	RemoveFriend(ctx context.Context, req *RemoveFriendDTO) error
}

func NewUsecase(deps Deps) *usecase {
	return &usecase{
		Deps: deps,
	}
}

type (
	TxManager interface {
		RunReadCommitted(ctx context.Context, f func(txCtx context.Context) error) error
	}
)

type SubscriberRepository interface {
	AcceptRequest(ctx context.Context, filter *subscriber_models.FriendRequestUpdateFilter) error
	MakeFriendRequest(ctx context.Context, reqfriendRequest *models.FriendRequest) error
	RejectRequest(ctx context.Context, filter *subscriber_models.FriendRequestUpdateFilter) error
	RemoveFriendRequest(ctx context.Context, filters *subscriber_models.FriendRequestDeleteFilter) error
}

type Deps struct {
	// Adapters
	SubscriberRepository
	//TxManager TxManager
}

var (
	_ Usecase = (*usecase)(nil)
)

type usecase struct {
	Deps
}
