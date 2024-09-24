package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

// @Summary		Health check application
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
