package subscriber

import "github.com/igortoigildin/go-contacts/subscriber/internal/app/models"

type (
	FriendRequestSearchFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
	FriendRequestInsertFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
	FriendRequestUpdateFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
	FriendRequestDeleteFilter struct {
		SenderIDs   []models.SenderID
		ReceiverIDs []models.ReceiverID
		Statuses    string
	}
)

func NewFriendRequestSearchFilter() *FriendRequestSearchFilter {
	return new(FriendRequestSearchFilter)
}

func NewFriendInsertFilter() *FriendRequestInsertFilter {
	return new(FriendRequestInsertFilter)
}

func NewFriendUpdateFilter() *FriendRequestUpdateFilter {
	return new(FriendRequestUpdateFilter)
}

func NewFriendDeleteFilter() *FriendRequestDeleteFilter {
	return new(FriendRequestDeleteFilter)
}
