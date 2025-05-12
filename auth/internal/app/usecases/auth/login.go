package auth

import "context"

func (u *usecase) Login(ctx context.Context, userInfo *LoginDTO) (token string, err error) {
	user, err := u.UserService.GetUserByUsername(ctx, userInfo.Username)
	if err != nil {
		return "", err
	}

	//TODO: implement login logic

	return user.Username, nil
}
