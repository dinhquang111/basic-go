package routes

import (
	"test-go/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/", handlers.HandleTelegramBotMessage)
	router.POST("/health", handlers.HealthCheck)
	AddSearchRoutes(router)
}
