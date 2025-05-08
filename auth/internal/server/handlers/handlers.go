package handlers

import (
	"net/http"
	"strings"

	models "github.com/igortoigildin/go-contacts/auth/internal/server/models"
	"github.com/labstack/echo/v4"
)

func RegisterHandler(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request"})
	}

	// TODO: Save user to DB
	// if user exists -> return 400

	return c.JSON(http.StatusCreated, models.SuccessResponse{Message: "User registered successfully"})
}

func LoginHandler(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request"})
	}

	// TODO: Validate user credentials
	var admin string = "admin"
	if req.Username != &admin || req.Password != &admin {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid credentials"})
	}

	// TODO: Generate JWT
	token := "mocked.jwt.token"

	return c.JSON(http.StatusOK, models.LoginResponse{Token: token})
}

func LogoutHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Missing token"})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid token format"})
	}

	// TODO: Invalidate token

	return c.JSON(http.StatusOK, models.SuccessResponse{Message: "Logged out successfully"})
}

func VerifyHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Missing token"})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid token format"})
	}

	// TODO: Verify token
	username := "admin"

	return c.JSON(http.StatusOK, models.VerifyResponse{
		Valid:    true,
		Username: username,
	})
}
