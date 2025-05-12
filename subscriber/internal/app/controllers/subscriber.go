package controllers

import (
	"context"

	pb "github.com/igortoigildin/go-contacts/subscriber/pkg/proto"
)

func (c *Controller) MakeFriendRequest(ctx context.Context, req *pb.FriendRequest) error {
	return nil
}

func (c *Controller) AcceptFriendRequest(ctx context.Context, req *pb.FriendAcceptRequest) error {
	return nil
}

func (c *Controller) RejectFriendRequest(ctx context.Context, req *pb.RejectFriendRequest) error {
	return nil
}

func (c *Controller) RemoveFriend(ctx context.Context, req *pb.RemoveFriendRequest) error {
	return nil
}

func (c *Controller) ListFriends(ctx context.Context, req *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	return nil, nil
}
