package subscriberrepository

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func convertPGError(err error) error {
	if err == nil {
		return nil
	}
	// https://github.com/jackc/pgx/wiki/Error-Handling
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Println(pgErr.Message) // => syntax error at end of input
		log.Println(pgErr.Code)    // => 42601

		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return fmt.Errorf("%s: %w", pgErr.Message, models.ErrAlreadyExists)
		default:
			return err
		}
	}
	return err
}

func NewFriendRequest(r *models.FriendRequest) (*friendRequest, error) {
	requestID, err := uuid.Parse(r.ID.String())
	if err != nil {
		return nil, err
	}

	receiverID, err := uuid.Parse(r.ReceiverID.String())
	if err != nil {
		return nil, err
	}

	senderID, err := uuid.Parse(r.SenderID.String())
	if err != nil {
		return nil, err
	}

	return &friendRequest{
		ID:         requestID,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Status:     r.Status,
		CreatedAt:  r.CreatedAt,
	}, nil
}
