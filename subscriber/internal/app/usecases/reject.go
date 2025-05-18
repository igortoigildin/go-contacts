package usecases

import (
	"context"
	"fmt"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
	usecase_models "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *usecase) RejectFriendRequest(ctx context.Context, req *FriendRejectDTO) (*emptypb.Empty, error) {
	filter := usecase_models.NewFriendUpdateFilter()
	filter.ReceiverIDs = append(filter.ReceiverIDs, models.ReceiverID(req.TargetUsername))
	filter.SenderIDs = append(filter.SenderIDs, models.SenderID(req.Username))

	err := u.SubscriberRepository.RejectRequest(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return nil, err
}
