package controllers

import (
	"context"

	pb "github.com/igortoigildin/go-contacts/subscriber/pkg/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Controller) MakeFriendRequest(ctx context.Context, req *pb.FriendRequest) (*emptypb.Empty, error) {
	// 1. convert delivery models to domain models/DTO
	subscriberInfo := convertProtoToDomain(req)

	// 2. call usecase
	_, err := c.SubscriberUsecase.MakeFriendRequest(ctx, subscriberInfo)

	return nil, err
}

func (c *Controller) AcceptFriendRequest(ctx context.Context, req *pb.FriendAcceptRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (c *Controller) RejectFriendRequest(ctx context.Context, req *pb.RejectFriendRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (c *Controller) RemoveFriend(ctx context.Context, req *pb.RemoveFriendRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (c *Controller) ListFriends(ctx context.Context, req *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	return nil, nil
}
