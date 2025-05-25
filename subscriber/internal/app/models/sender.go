package models

type SenderID string

type Sender struct {
	ID SenderID
}

func (id SenderID) String() string {
	return string(id)
}

func (s *Sender) String() string {
	return string(s.ID)
}
