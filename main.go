package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong " + c.Params.ByName("name"),
		})
	})

	// Run the server
	router.Run(":8080")
}
