package middleware

import (
	"auth-api/models"
	"auth-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnUnauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error: []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeUnauthorized,
				ErrorMessage: "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
			return
		}

		// Vai deixar somente o token removendo a palavra Bearer
		token := header[len(Bearer_schema):]

		valid, claims := services.VerifyToken(token)

		if !valid {
			ReturnUnauthorized(c)
		}

		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}

		c.Keys["Sum"] = claims.Sum
		c.Keys["Roles"] = claims.Roles
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(context *gin.Context) {

		if len(context.Keys) == 0 {
			ReturnUnauthorized(context)
		}

		rolesVal := context.Keys["Roles"]
		fmt.Println("roles", rolesVal)
		if rolesVal == nil {
			ReturnUnauthorized(context)
		}

		roles := rolesVal.([]string)
		validation := make(map[string]string)
		for _, val := range roles {
			validation[val] = ""
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				ReturnUnauthorized(context)
			}
		}
	}
}
