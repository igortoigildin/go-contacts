package user

import (
	"net/http"

	models "github.com/igortoigildin/go-contacts/user/internal/models"
	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {
	var user models.User

	_, err := CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func UpdateUserHandler(c echo.Context) error {
	id := c.Param("id")

	_, err := UpdateUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}

func GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	user, err := GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func SearchUserHandler(c echo.Context) error {
	nickname := c.QueryParam("nickname")
	result := SearchUsers(nickname)
	return c.JSON(http.StatusOK, result)
}

func CreateUser(user models.User) (any, error) {
	// TODO: to be updated
	return user, nil
}

func UpdateUser(id string) (any, error) {
	// TODO: to be updated
	var user models.User
	return user, nil
}

func GetUser(id string) (any, error) {
	// TODO: to be updated
	var user models.User
	return user, nil
}

func SearchUsers(id string) any {
	// TODO: to be updated
	var user models.User
	return user
}
