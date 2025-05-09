package controllers

import (
	"github.com/igortoigildin/go-contacts/auth/internal/app/usecases/auth"

	pb "github.com/igortoigildin/go-contacts/auth/pkg/proto"
)


func newRegisterDTOFromRegisterRequest(req *pb.RegisterRequest) *auth.RegistrationDTO {
	return &auth.RegistrationDTO{
		Username: req.Username,
		Password: req.Password,
	}
}

func newLoginDTOFromLoginRequest(req *pb.LoginRequest) *auth.LoginDTO {	
	return &auth.LoginDTO{
		Username: req.Username,
		Password: req.Password,
	}
}

func newLogoutDTOFromLogoutRequest(req *pb.LogoutRequest) *auth.LogoutDTO {
	return &auth.LogoutDTO{
		Username: req.Username,
	}
}

func newVerifyDTOFromVerifyRequest(req *pb.VerifyRequest) *auth.VerifyDTO {
	return &auth.VerifyDTO{
		Username: req.Username,
		Token: req.Token,
	}
}





