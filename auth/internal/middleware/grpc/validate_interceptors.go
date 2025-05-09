package grpc_middleware

import (
	"context"

	"buf.build/go/protovalidate"
	grpcutils "github.com/igortoigildin/go-contacts/auth/pkg/grpc_utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// ValidateUnaryServerInterceptor - convert any arror to rpc error
func ValidateUnaryServerInterceptor(validator protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		if msg, ok := (req).(proto.Message); ok {
			if err := validator.Validate(msg); err != nil {
				return nil, grpcutils.RPCValidationError(err)
			}
		}

		return handler(ctx, req)
	}
}
