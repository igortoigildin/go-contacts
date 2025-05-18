package main

import (
	"context"
	"log"
	"time"

	"buf.build/go/protovalidate"
	"github.com/igortoigildin/go-contacts/subscriber/pkg/postgres"
	"github.com/igortoigildin/go-contacts/subscriber/pkg/postgres/transaction_manager"
	"google.golang.org/grpc"

	subscriber_controller "github.com/igortoigildin/go-contacts/subscriber/internal/app/controllers"
	subscriber_repository "github.com/igortoigildin/go-contacts/subscriber/internal/app/repositories/subscriber/postgres"
	"github.com/igortoigildin/go-contacts/subscriber/internal/app/server"
	subscriber_usecase "github.com/igortoigildin/go-contacts/subscriber/internal/app/usecases"
	grpc_middleware "github.com/igortoigildin/go-contacts/subscriber/internal/middleware/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// =========================
	// adapters
	// =========================

	// repository
	pgConn, err := postgres.NewConnectionPool(ctx, "postgres://jack:secret@pg.example.com:5432/mydb?sslmode=disabled",
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}

	txManager := transaction_manager.New(pgConn)

	subscriberRepo := subscriber_repository.NewRepository(txManager)

	// =========================
	// usecases
	// =========================
	subscriberUsecase := subscriber_usecase.NewUsecase(subscriber_usecase.Deps{
		SubscriberRepository: subscriberRepo,
	})

	// =========================
	// delivery
	// =========================
	subscriberController := subscriber_controller.New(subscriber_controller.Deps{
		SubscriberUsecase: subscriberUsecase,
	})

	// middlewares
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
		ChainUnaryInterceptors: mws,
	}

	srv, err := server.New(ctx, config, server.Controllers{
		SubscriberServiceServer: subscriberController,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
