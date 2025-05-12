package app

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	chat "github.com/igortoigildin/go-contacts/chat/internal/api"
)

// TODO: to be updated. Please refer to USER service

type App struct {}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	e := echo.New()

	// Chat routes
	e.POST("/chats", chat.CreateChatHandler)
	e.GET("/chats/:userID", chat.GetUserChatsHandler)
	e.POST("/chats/:chatID/messages", chat.SendMessageHandler)
	e.GET("/chats/:chatID/messages", chat.GetMessagesHandler)

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