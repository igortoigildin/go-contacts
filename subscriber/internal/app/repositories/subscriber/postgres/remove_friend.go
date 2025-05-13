package subscriberrepository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	request_models "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/models"
	usecasemodels "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/models"
)

func (r *Repository) RemoveFriendRequest(ctx context.Context, filters *usecasemodels.FriendRequestSearchFilter) error {

	qb := squirrel.Delete(tableFriendRequests).
		PlaceholderFormat(squirrel.Dollar)

	qb = applyRequestsDeleteFilter(qb, filters)

	conn := r.txManager.GetQueryEngine(ctx)
	if _, err := conn.Execx(ctx, qb); err != nil {
		return fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return nil
}

func applyRequestsDeleteFilter(qb squirrel.DeleteBuilder, filter *request_models.FriendRequestSearchFilter) squirrel.DeleteBuilder {
	if filter == nil {
		return qb
	}

	if len(filter.SenderIDs) > 0 {
		qb = qb.Where(squirrel.Eq{columnSenderID: filter.SenderIDs})
	}

	if len(filter.ReceiverIDs) > 0 {
		qb = qb.Where(squirrel.Eq{columnReceiverID: filter.ReceiverIDs})
	}

	if len(filter.Statuses) > 0 {
		qb = qb.Where(squirrel.Eq{columnStatus: filter.Statuses})
	}

	return qb
}
