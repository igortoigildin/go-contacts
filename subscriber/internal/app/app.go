package app

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	subscriber "github.com/igortoigildin/go-contacts/subscriber/internal/api"
)

// TODO: to be updated. Please refer to USER service
type App struct {}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
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

	return nil
}