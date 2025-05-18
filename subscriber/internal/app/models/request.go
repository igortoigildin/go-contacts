package models

import (
	"time"

	"github.com/google/uuid"
)

// FriendRequestID is the ID of a friend request
type FriendRequestID string

func (id FriendRequestID) String() string {
	return string(id)
}

type FriendRequest struct {
	ID         FriendRequestID
	SenderID   SenderID
	ReceiverID ReceiverID
	Status     string
	CreatedAt  time.Time
}

func NewFriendRequest() *FriendRequest {
	return &FriendRequest{
		ID:        FriendRequestID(uuid.New().String()),
		Status:    "pending",
		CreatedAt: time.Now(),
	}
}
