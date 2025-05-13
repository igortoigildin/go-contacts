package usecases

import (
	"context"
	"fmt"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
)

func (u *usecase) RejectFriendRequest(ctx context.Context, req *FriendRejectDTO) error {
	filter := NewFriendUpdateFilter()
	filter.ReceiverIDs = append(filter.ReceiverIDs, models.ReceiverID(req.TargetUsername))
	filter.SenderIDs = append(filter.SenderIDs, models.SenderID(req.Username))

	// Save reject request to DB
	err := u.TxManager.RunReadCommitted(ctx,

		func(txCtx context.Context) error {

			if err := u.SubscriberRepository.RejectRequest(txCtx, filter); err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		},
	)

	return err
}
