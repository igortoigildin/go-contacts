package subscriberrepository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
)

func (r *Repository) MakeFriendRequest(ctx context.Context, reqfriendRequest *models.FriendRequest) error {
	row, err := NewFriendRequest(reqfriendRequest)
	if err != nil {
		return err
	}

	qb := squirrel.Insert(tableFriendRequests).
		Columns(friendRequestsColumns...).
		Values(row.Values(friendRequestsColumns...)...).
		PlaceholderFormat(squirrel.Dollar)

	conn := r.txManager.GetQueryEngine(ctx)
	if _, err = conn.Execx(ctx, qb); err != nil {
		return fmt.Errorf("postgres: %w", convertPGError(err))
	}

	return nil
}
