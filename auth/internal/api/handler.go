package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	models "github.com/igortoigildin/go-contacts/auth/internal/models"
	token "github.com/igortoigildin/go-contacts/auth/pkg/utils"
)

func LoginHandler(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	userID, err := Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}
	token, _ := token.GenerateToken(userID)
	return c.JSON(http.StatusOK, models.AuthResponse{AccessToken: token})
}

func VerifyHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := token.ParseToken(tokenStr)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"user_id": userID})
}

func LogoutHandler(c echo.Context) error {
	// TODO: business logic to be added
	return c.JSON(http.StatusOK, map[string]string{"message": "logged out"})
}

func Login(email, password string) (string, error) {
	// TODO: business logic to be added
	return "1", nil
}
