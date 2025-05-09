package controllers

import (
	"buf.build/go/protovalidate"

	pb "github.com/igortoigildin/go-contacts/auth/pkg/proto"

	auth "github.com/igortoigildin/go-contacts/auth/internal/app/usecases/auth"
)

type Controller struct {
	pb.UnimplementedAuthServiceServer

	Validator protovalidate.Validator

	Deps
}

func New(deps Deps) *Controller {
	return &Controller{
		Deps: deps,
	}
}

type Deps struct {
	AuthUsecase auth.Usecase
}