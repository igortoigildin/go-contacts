package subscriber

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeFriendRequestHandler(c echo.Context) error {
	err := makeRequest("1", "2")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "friend request sent"})
}

func AcceptFriendHandler(c echo.Context) error {
	err := acceptRequest("1", "2")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "friend request accepted"})
}

func RejectFriendHandler(c echo.Context) error {
	err := rejectRequest("1", "2")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "friend request rejected"})
}

func RemoveFriendHandler(c echo.Context) error {
	err := removeFriend("1", "2")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "friend removed"})
}

func GetFriendsHandler(c echo.Context) error {
	id := c.Param("id")
	friends, _ := getFriends(id)
	return c.JSON(http.StatusOK, echo.Map{"friends": friends})
}

// TODO: business logic to be updated
func makeRequest(idFrom, idTo string) error {
	return nil
}

func acceptRequest(idFrom, idTo string) error {
	return nil
}

func rejectRequest(idFrom, idTo string) error {
	return nil
}

func removeFriend(idFrom, idTo string) error {
	return nil
}

func getFriends(id string) (any, error) {
	return nil, nil
}
