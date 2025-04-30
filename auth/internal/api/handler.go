package api

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	models "github.com/igortoigildin/go-contacts/auth/internal/models"
	token "github.com/igortoigildin/go-contacts/auth/pkg/utils"
)

func RegisterHandler(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "username and password required"})
	}

	userID := uuid.New().String()

	_, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "hashing failed"})
	}

	token, err := token.GenerateToken(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "token generation failed"})
	}

	// TODO: Save user in DB
	// TODO: publish event in message broker

	return c.JSON(http.StatusOK, models.AuthResponse{
		Token:  token,
		UserID: userID,
	})
}

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
	return c.JSON(http.StatusOK, models.AuthResponse{Token: token})
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
