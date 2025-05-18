package usecases

import (
	"context"
	"fmt"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *usecase) MakeFriendRequest(ctx context.Context, req *FriendRequestDTO) (*emptypb.Empty, error) {
	friendRequest := models.NewFriendRequest()
	friendRequest.ReceiverID = models.ReceiverID(req.TargetUsername)
	friendRequest.SenderID = models.SenderID(req.Username)

	err := u.SubscriberRepository.MakeFriendRequest(ctx, friendRequest)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return nil, err
}
