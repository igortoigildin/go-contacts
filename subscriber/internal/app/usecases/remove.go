package usecases

import (
	"context"
	"fmt"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
	"google.golang.org/protobuf/types/known/emptypb"

	usecase_models "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/models"
)

func (u *usecase) RemoveFriend(ctx context.Context, req *RemoveFriendDTO) (*emptypb.Empty, error) {
	filter := usecase_models.NewFriendDeleteFilter()
	filter.ReceiverIDs = append(filter.ReceiverIDs, models.ReceiverID(req.FriendUsername))
	filter.SenderIDs = append(filter.SenderIDs, models.SenderID(req.Username))

	err := u.SubscriberRepository.RemoveFriendRequest(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return nil, err
}
