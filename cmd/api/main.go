package main

import (
	_ "basic-go/docs"
	"basic-go/internal/api/routes"
	"basic-go/internal/jenkins"
	"basic-go/internal/logger"
	"basic-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// @title           Service API
// @version         1.0

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host      localhost:8080
// @BasePath  /

func main() {
	gin.SetMode(gin.ReleaseMode)
	jenkins.SetupJenkin()
	logger.SetupLogger()
	engine := gin.New()
	engine.Use(middleware.Logger())
	engine.Use(middleware.ErrorHandler())
	routes.SetupRoutes(engine)
	engine.Run(":8080")
}
