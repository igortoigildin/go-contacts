package auth

import (
	"context"
)

type Usecase interface {
	Register(ctx context.Context, userInfo *RegistrationDTO) (username string, err error)
	Login(ctx context.Context, userInfo *LoginDTO) (token string, err error)
	Logout(ctx context.Context, userInfo *LogoutDTO) (username string, err error)
	Verify(ctx context.Context, userInfo *VerifyDTO) (username string, err error)
}


var (
	_ Usecase = (*usecase)(nil)
)

type usecase struct {
	Deps
}

// business logic
func NewUsecase(deps Deps) *usecase {
	return &usecase{
		Deps: deps,
	}
}

type Deps struct {
	// UserRepository repository.UserRepository
}
