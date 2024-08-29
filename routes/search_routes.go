package routes

import (
	handlesearch "test-go/handlers"

	"github.com/gin-gonic/gin"
)

func UseSearchRoutes(router *gin.Engine) {
	searchGroup := router.Group("/search")
	{
		searchGroup.GET("/", handlesearch.Query)
		searchGroup.POST("/", handlesearch.Insert)
	}
}
