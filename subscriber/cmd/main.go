package main

import (
	"log"

	"github.com/igortoigildin/go-contacts/subscriber/internal/app"
)


func main() {
	app := app.New()
	if err := app.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
