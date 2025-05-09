package main

import (
	"context"
	"log"

	"github.com/igortoigildin/go-contacts/auth/internal/app/server"
)


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// infrastructure server
	config := server.Config{
		GRPCPort:               ":8082",
		GRPCGatewayPort:        ":8080",
		ChainUnaryInterceptors: mws,
	}

	srv, err := server.New(ctx, config, server.Contollers{
		OrdersManagementSystemServiceServer: ordersController,
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err = srv.Run(ctx); err != nil {
		log.Fatalf("run: %v", err)
	}
}
