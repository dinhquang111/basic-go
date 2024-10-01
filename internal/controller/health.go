package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

// HealthCheck @Summary		Health check application
// @Description	Health check and return app's metadata
// @Produce		json
// @Success		200	{object}	string
// @Router			/health [get]
func (c *Controller) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version":   Version,
		"commit":    Commit,
		"buildTime": BuildTime,
	})
}

// HealthCheck1 @Summary		Health check application
// @Description	Health check and return app's metadata
// @Produce		json
// @Success		200	{object}	string
// @Router			/health [post]
func (c *Controller) HealthCheck1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version":   Version,
		"commit":    Commit,
		"buildTime": BuildTime,
	})
}
