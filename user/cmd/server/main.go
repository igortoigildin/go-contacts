package main

import (
	"context"
	"log"

	"github.com/igortoigildin/go-contacts/user/internal/server/app"
)

func main() {
	app, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	log.Println("app initialized")

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
