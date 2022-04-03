package handlers

import (
	"auth-api/commands"
	"auth-api/domain/entities"
	"auth-api/infra/database"
	"auth-api/models"
	"auth-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoginHandler(c *gin.Context) {
	db := database.GetGatabase()

	var login commands.LoginCommand
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user entities.User
	dbError := db.Where("email = ?", login.Email).First(&user).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"message": "user already exists",
		})

		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(401, gin.H{
			"message": "Invalid credentials",
		})

		return
	}

	var claims = &services.Claim{}
	claims.Sum = user.ID
	claims.Roles = []string{"Manager", "User", "Admin"} // Pegar as roles do banco
	//claims.Audience = c.Request.Header.Get("Referer")

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2 * time.Hour))
	token, err := services.NewJwtService().GenerateToken(claims, expirationTime)
	if err != nil {
		badRequest(c, http.StatusBadRequest, "error in gerating token", []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: err.Error(),
			},
		})
	}

	ok(c, http.StatusOK, "Token created!", token)
}
