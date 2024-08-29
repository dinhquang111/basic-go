package routes

import "github.com/gin-gonic/gin"

func UseRoutes(router *gin.Engine) {
	UseSearchRoutes(router)
}
