package models

type ReceiverID string

func (id ReceiverID) String() string {
	return string(id)
}

type Receiver struct {
	ID ReceiverID
}
