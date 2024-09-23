package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		raw := c.Request.URL.RawQuery
		id := uuid.NewString()

		log.Info().
			Str("id", id).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Msg(raw)

		c.Next()

		var zeroLog *zerolog.Event
		latency := time.Since(start).Milliseconds()
		status := c.Writer.Status()
		switch {
		case status < 400:
			zeroLog = log.Info()
		case status < 500:
			zeroLog = log.Warn()
		default:
			zeroLog = log.Error()
		}
		zeroLog.
			Str("id", id).
			Int("status", c.Writer.Status()).
			Int64("latency", latency).
			Msg("HTTP request completed")
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := c.Writer.Status()
		if status < 500 {
			return
		}
		c.Errors = c.Errors[:0]
		c.JSON(500, gin.H{
			"error": "An internal server error occurred",
		})
		c.Abort()
	}
}
