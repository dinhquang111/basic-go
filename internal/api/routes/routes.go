package routes

import (
	"test-go/internal/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {
	c := controller.NewController()
	// router.POST("/", handlers.HandleTelegramBotMessage)
	router.GET("/health", c.HealthCheck)
	AddSearchRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
