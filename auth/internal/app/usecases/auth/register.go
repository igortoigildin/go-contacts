package auth

import (
	"context"
	"errors"
)

var (
	ErrRegisterFailed = errors.New("register failed")
)

func (uc *usecase) Register(ctx context.Context, userInfo *RegistrationDTO) (username string, err error) {
	user, err := uc.UserService.GetUserByUsername(ctx, userInfo.Username)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
