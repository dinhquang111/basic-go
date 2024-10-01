package handlers

import (
	"basic-go/internal/jenkins"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TelegramMessagePayload struct {
	Message Message `json:"message"`
}

type Message struct {
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type User struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int64 `json:"id"`
}
type RequestMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func ValidateJenkinsJob() {
	job, err := jenkins.Client.GetJob(jenkins.Ctx, "Backend", "CLS", "Staging")
	if err != nil {
		panic(err)
	}
	details, _ := job.GetParameters(jenkins.Ctx)
	isBuilding, _ := job.IsRunning(jenkins.Ctx)
	fmt.Print(details, isBuilding)
}

func HandleTelegramBotMessage(c *gin.Context) {
	var payload TelegramMessagePayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// triggerUsername := payload.Message.From.Username
	// messages := strings.Split(payload.Message.Text, string("\n"))
	// for _, message := range messages {

	// }

	// data := RequestMessage{
	// 	ChatID: payload.Message.Chat.ID,
	// 	Text:   message + ` build is triggered by ` + triggerUserId + `.`,
	// }

	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	return
	// }
	// response, err := http.Post("https://api.telegram.org/bot6654034396:AAEc3hoa3r11NRfMb9ALhXmjWjNzOvozEds/sendMessage", "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer response.Body.Close()

	c.JSON(http.StatusOK, gin.H{})
}
