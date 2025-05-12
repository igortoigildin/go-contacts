package main

import (
	"context"
	"log"

	"buf.build/go/protovalidate"
	"github.com/igortoigildin/go-contacts/auth/internal/app/server"
	"github.com/igortoigildin/go-contacts/auth/internal/app/usecases/auth"
	"google.golang.org/grpc"

	userservice "github.com/igortoigildin/go-contacts/auth/internal/app/adapters/user_service"
	authController "github.com/igortoigildin/go-contacts/auth/internal/app/controllers"
	grpc_middleware "github.com/igortoigildin/go-contacts/auth/internal/middleware/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// services
	userServiceAdapter := userservice.NewClient()

	// =========================
	// usecases
	// =========================

	authUsecase := auth.NewUsecase(auth.Deps{
		UserService: userServiceAdapter,
	})

	// controllers
	authController := authController.New(authController.Deps{
		AuthUsecase: authUsecase,
	})

	// middlewares
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("server: failed to initialize validator: %s", err)
	}
	mws := []grpc.UnaryServerInterceptor{
		grpc_middleware.ErrorsUnaryServerInterceptor(),
		grpc_middleware.ValidateUnaryServerInterceptor(validator),
	}

	// infrastructure server
	config := server.Config{
		GRPCPort:               ":8082",
		GRPCGatewayPort:        ":8080",
		ChainUnaryInterceptors: mws,
	}

	srv, err := server.New(ctx, config, server.Controllers{
		AuthServiceServer: authController,
	},
	)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
