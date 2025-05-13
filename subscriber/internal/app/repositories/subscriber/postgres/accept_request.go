package subscriberrepository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	request_models "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases/models"
)

func (r *Repository) AcceptRequest(ctx context.Context, filter *request_models.FriendRequestUpdateFilter) error {
	qb := squirrel.Update(tableFriendRequests).Set(columnStatus, "accepted").
		Set(columnCreatedAt, squirrel.Expr("CURRENT_TIMESTAMP")).
		PlaceholderFormat(squirrel.Dollar)

	qb = applyRequestsUpdateFilter(qb, filter)

	conn := r.txManager.GetQueryEngine(ctx)
	if _, err := conn.Execx(ctx, qb); err != nil {
		return fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return nil
}

func applyRequestsUpdateFilter(qb squirrel.UpdateBuilder, filter *request_models.FriendRequestUpdateFilter) squirrel.UpdateBuilder {
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
