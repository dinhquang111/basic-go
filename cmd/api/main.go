package main

import (
	_ "basic-go/docs"
	"basic-go/internal/api/routes"
	"basic-go/internal/jenkins"
	"basic-go/internal/logger"
	"basic-go/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// @title           Service API
// @version         1.0

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.ioT

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath  /

func main() {
	jenkins.SetupJenkins()
	logger.SetupLogger()
	engine := gin.New()
	engine.Use(middleware.Logger())
	engine.Use(middleware.ErrorHandler())
	routes.SetupRoutes(engine)
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal().Msg("Failed to start server")
	}
}
