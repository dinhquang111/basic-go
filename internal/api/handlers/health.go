package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version":   Version,
		"commit":    Commit,
		"buildTime": BuildTime,
	})
}

type TelegramMessagePayload struct {
	UpdateID int64   `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int64  `json:"message_id"`
	From      User   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
type RequestMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func HandleTelegramBotMessage(c *gin.Context) {
	var payload TelegramMessagePayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	triggerUserId := payload.Message.From.Username
	message := payload.Message.Text

	data := RequestMessage{
		ChatID: payload.Message.Chat.ID,
		Text:   message + ` build is triggered by ` + triggerUserId + `.`,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	response, err := http.Post("https://api.telegram.org/bot6654034396:AAEc3hoa3r11NRfMb9ALhXmjWjNzOvozEds/sendMessage", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	c.JSON(http.StatusOK, gin.H{})
}
