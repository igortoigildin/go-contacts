package main

import (
	"context"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	desc "github.com/igortoigildin/go-contacts/user/pkg/api/users"
	grpc_utils "github.com/igortoigildin/go-contacts/user/pkg/grpc_utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// Define retry options
	opts := []grpc_retry.CallOption{
		grpc_retry.WithMax(3), // max 3 retry attempts
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithCodes(codes.Unavailable, codes.DeadlineExceeded),
	}

	conn, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	userClient := desc.NewUserServiceClient(conn)

	// init Circuit Breaker

	cb := grpc_utils.InitCircuitBreaker()

	// /CreateUser
	{
		result, err := cb.Execute(func() (interface{}, error) {
			// call with breaker
			resp, err := userClient.CreateUser(context.Background(), &desc.CreateUserRequest{
				Email: "test@test.com",
				Name:  "test",
			})
			if err != nil {
				log.Fatalf("failed to create user: %v", err)
			}

			return resp, nil
		})
		resp := result.(*desc.User)

		user, err := protojson.Marshal(resp)
		if err != nil {
			log.Fatalf("failed to marshal user: %v", err)
		}
		log.Printf("created user: %v", string(user))
	}

	// /GetUser
	{
		result, err := cb.Execute(func() (interface{}, error) {
			// call with breaker
			resp, err := userClient.GetUser(context.Background(), &desc.GetUserRequest{
				Id: "12345",
			})
			if err != nil {
				log.Fatalf("failed to get user: %v", err)
			}
			return resp, nil
		})
		resp := result.(*desc.User)
		user, err := protojson.Marshal(resp)
		if err != nil {
			log.Fatalf("failed to marshal user: %v", err)
		}
		log.Printf("got user: %v", string(user))
	}

	// /UpdateUser
	{
		result, err := cb.Execute(func() (interface{}, error) {
			// call with breaker
			resp, err := userClient.UpdateUser(context.Background(), &desc.UpdateUserRequest{
				Id:    "12345",
				Email: "test@test.com",
			})
			if err != nil {
				log.Fatalf("failed to update user: %v", err)
			}
			return resp, nil
		})
		resp := result.(*desc.User)
		user, err := protojson.Marshal(resp)
		if err != nil {
			log.Fatalf("failed to marshal user: %v", err)
		}
		log.Printf("updated user: %v", string(user))
	}

	// /SearchUser
	{
		result, err := cb.Execute(func() (interface{}, error) {
			// call with breaker
			resp, err := userClient.SearchUsers(context.Background(), &desc.SearchUsersRequest{
				Query: "test",
			})
			if err != nil {
				log.Fatalf("failed to search user: %v", err)
			}
			return resp, nil
		})
		resp := result.(*desc.SearchUsersResponse)

		user, err := protojson.Marshal(resp)
		if err != nil {
			log.Fatalf("failed to marshal user: %v", err)
		}
		log.Printf("searched user: %v", string(user))
	}
}
