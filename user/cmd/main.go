package main

import (
	"net/http"
	"os"
	"time"

	user "github.com/igortoigildin/go-contacts/user/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/users", user.CreateUserHandler)
	e.PUT("/users/:id", user.UpdateUserHandler)
	e.GET("/users/:id", user.GetUserHandler)
	e.GET("/users/search", user.SearchUserHandler)

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
