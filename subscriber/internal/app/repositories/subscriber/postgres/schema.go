package subscriberrepository

import (
	"time"

	"github.com/google/uuid"
)

const tableFriendRequests = "public.friend_requests"

const (
	columnID         = "request_id"
	columnSenderID   = "sender_id"
	columnReceiverID = "receiver_id"
	columnStatus     = "status"
	columnCreatedAt  = "created_at"
)

type friendRequest struct {
	ID         uuid.UUID `db:"request_id"`
	SenderID   uuid.UUID `db:"sender_id"`
	ReceiverID uuid.UUID `db:"receiver_id"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
}

var (
	friendRequestsColumns = []string{
		columnID,
		columnSenderID,
		columnReceiverID,
		columnStatus,
		columnCreatedAt,
	}
)

func (r *friendRequest) mapFields() map[string]any {
	return map[string]any{
		columnID:         r.ID,
		columnSenderID:   r.SenderID,
		columnReceiverID: r.ReceiverID,
	}
}

func (r *friendRequest) Values(columns ...string) []any {
	mapFields := r.mapFields()
	values := make([]any, 0, len(columns))
	for _, column := range columns {
		if v, ok := mapFields[column]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}
