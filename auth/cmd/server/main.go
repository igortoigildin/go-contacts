package main

import (
	"net/http"
	"os"
	"time"

	auth "github.com/igortoigildin/go-contacts/auth/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/auth/register", auth.RegisterHandler)
	e.POST("/auth/login", auth.LoginHandler)
	e.POST("/auth/logout", auth.LogoutHandler)
	e.GET("/auth/verify", auth.VerifyHandler)

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
