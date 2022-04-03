package routes

import (
	"auth-api/handlers"
	"auth-api/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		main.POST("/login", handlers.LoginHandler)

		user := main.Group("user").Use(middleware.ValidateToken()).Use(middleware.Authorization([]string{"Managere"}))
		{
			user.POST("/", handlers.CreateUserHandler)
		}

	}
	return router
}
