package controllers

import (
	usecase "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases"
	pb "github.com/igortoigildin/go-contacts/subscriber/pkg/proto"
)

func convertProtoToDomain(req *pb.FriendRequest) *usecase.FriendRequestDTO {
	return &usecase.FriendRequestDTO{
		Username:       req.GetUsername(),
		TargetUsername: req.GetTargetUsername(),
	}
}
