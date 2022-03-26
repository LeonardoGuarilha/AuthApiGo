package handlers

import (
	"auth-api/domain/entities"
	"auth-api/infra/database"
	"auth-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	db := database.GetGatabase()

	var user entities.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "can't create user: " + err.Error(),
		})

		return
	}

	c.JSON(201, user)
}
