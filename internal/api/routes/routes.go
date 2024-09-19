package routes

import (
	"test-go/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", handlers.HealthCheck)
	AddSearchRoutes(router)
}
