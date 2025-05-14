package usecases

import (
	"context"
	"fmt"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
)

func (u *usecase) MakeFriendRequest(ctx context.Context, req *FriendRequestDTO) error {
	// Create new friend request
	friendRequest := models.NewFriendRequest()
	friendRequest.ReceiverID = models.ReceiverID(req.TargetUsername)
	friendRequest.SenderID = models.SenderID(req.Username)

	// // Save request to DB

	err := u.SubscriberRepository.MakeFriendRequest(ctx, friendRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return err
}
