package handlers

import (
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
