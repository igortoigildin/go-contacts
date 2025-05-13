package usecases

import (
	"context"
	"fmt"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
)

func (u *usecase) RemoveFriend(ctx context.Context, req *RemoveFriendDTO) error {
	filter := NewFriendDeleteFilter()
	filter.ReceiverIDs = append(filter.ReceiverIDs, models.ReceiverID(req.FriendUsername))
	filter.SenderIDs = append(filter.SenderIDs, models.SenderID(req.Username))

	// Save update request to DB
	err := u.TxManager.RunReadCommitted(ctx,

		func(txCtx context.Context) error {

			if err := u.SubscriberRepository.RemoveFriendRequest(txCtx, filter); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	)

	return err
}
