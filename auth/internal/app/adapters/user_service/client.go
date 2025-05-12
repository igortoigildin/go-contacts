package userservice

import (
	"github.com/igortoigildin/go-contacts/auth/internal/app/usecases/auth"
)

var (
	_ auth.UserService = (*Client)(nil)
)

type Client struct {
	// TODO: add fields
}

func NewClient() *Client {
	return &Client{}
}
