package grpc_middleware

import (
	"context"
	"errors"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorsUnaryServerInterceptor - convert any arror to rpc error
func ErrorsUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			if _, ok := status.FromError(err); ok {
				return resp, err
			}

			switch {
			case errors.Is(err, models.ErrNotFound):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err, models.ErrAlreadyExists):
				return nil, status.Error(codes.AlreadyExists, err.Error())
			default:
				return nil, status.Error(codes.Internal, err.Error())
			}

		}

		return
	}
}
