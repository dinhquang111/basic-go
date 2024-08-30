package routes

import (
	"test-go/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func AddSearchRoutes(router *gin.Engine) {
	searchGroup := router.Group("/search")
	{
		searchGroup.GET("/", handlers.Query)
		searchGroup.POST("/", handlers.Insert)
	}
}
