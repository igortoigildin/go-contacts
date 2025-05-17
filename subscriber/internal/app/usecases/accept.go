package usecases

import (
	"context"
	"fmt"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
	usecase_models "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/models"
)

func (u *usecase) AcceptFriendRequest(ctx context.Context, req *FriendAcceptDTO) error {
	filter := usecase_models.NewFriendUpdateFilter()
	filter.ReceiverIDs = append(filter.ReceiverIDs, models.ReceiverID(req.TargetUsername))
	filter.SenderIDs = append(filter.SenderIDs, models.SenderID(req.Username))

	err := u.SubscriberRepository.AcceptRequest(ctx, filter)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return err
}
