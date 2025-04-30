package main

import (
	"net/http"
	"os"
	"time"

	subscriber "github.com/igortoigildin/go-contacts/subscriber/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Subscriber routes
	e.POST("/subscriber/request", subscriber.MakeFriendRequestHandler)
	e.POST("/subscriber/accept", subscriber.AcceptFriendHandler)
	e.POST("/subscriber/reject", subscriber.RejectFriendHandler)
	e.DELETE("/subscriber", subscriber.RemoveFriendHandler)
	e.GET("/subscriber/:id", subscriber.GetFriendsHandler)

	e.GET("/health/ready", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	time.Sleep(5 * time.Second)

	e.Logger.Fatal(e.Start(":" + httpPort))
}
