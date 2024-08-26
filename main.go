package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	version := 3
	router.GET("/ping/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong " + c.Params.ByName("name"),
			"version": "v" + strconv.Itoa(version),
		})
	})

	// Run the server
	router.Run(":8080")
}
