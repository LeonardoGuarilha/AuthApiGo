package routes

import (
	"auth-api/handlers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", handlers.CreateUserHandler)
		}
	}
	return router
}
