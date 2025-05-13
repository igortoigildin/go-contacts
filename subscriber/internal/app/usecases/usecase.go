package usecases

import (
	"context"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
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
	AcceptRequest(ctx context.Context, filter *FriendRequestUpdateFilter) error
	MakeFriendRequest(ctx context.Context, reqfriendRequest *models.FriendRequest) error
	RejectRequest(ctx context.Context, filter *FriendRequestUpdateFilter) error
	RemoveFriendRequest(ctx context.Context, filters *FriendRequestDeleteFilter) error
}

type Deps struct {
	// Adapters
	SubscriberRepository
	TxManager TxManager
}

var (
	_ Usecase = (*usecase)(nil)
)

type usecase struct {
	Deps
}



type (
	FriendRequestSearchFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
	FriendRequestInsertFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
	FriendRequestUpdateFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
	FriendRequestDeleteFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
)

func NewFriendRequestSearchFilter() *FriendRequestSearchFilter {
	return new(FriendRequestSearchFilter)
}

func NewFriendInsertFilter() *FriendRequestInsertFilter {
	return new(FriendRequestInsertFilter)
}

func NewFriendUpdateFilter() *FriendRequestUpdateFilter {
	return new(FriendRequestUpdateFilter)
}

func NewFriendDeleteFilter() *FriendRequestDeleteFilter {
	return new(FriendRequestDeleteFilter)
}
