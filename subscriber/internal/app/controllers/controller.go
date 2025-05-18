package controllers

import (
	"github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases"
	pb "github.com/igortoigildin/go-contacts/subscriber/pkg/proto"
)

type Controller struct {
	pb.UnimplementedSubscriberServiceServer

	Deps
}

func New(deps Deps) *Controller {
	return &Controller{
		Deps: deps,
	}
}

type Deps struct {
	SubscriberUsecase usecases.Usecase
}
