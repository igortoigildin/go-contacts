package chat

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateChatHandler(c echo.Context) error {
	chat, err := CreateChat("1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, chat)
}

func GetUserChatsHandler(c echo.Context) error {
	userID := c.Param("userID")
	chats, _ := GetChatsByUser(userID)
	return c.JSON(http.StatusOK, chats)
}

func SendMessageHandler(c echo.Context) error {
	chatID := c.Param("chatID")
	msg, err := SendMessage(chatID, "1", "2")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, msg)
}

func GetMessagesHandler(c echo.Context) error {
	chatID := c.Param("chatID")
	msgs, _ := GetMessages(chatID)
	return c.JSON(http.StatusOK, msgs)
}

// TODO: business logic to be updated
func CreateChat(id string) (any, error) {
	return nil, nil
}

func GetChatsByUser(id string) (any, error) {
	return nil, nil
}

func SendMessage(chatID, SenderID, Text string) (any, error) {
	return nil, nil
}

func GetMessages(id string) (any, error) {
	return nil, nil
}
