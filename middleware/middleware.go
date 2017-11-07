package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liteByte/mascotas-backend/authentication"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		token := authentication.GetTokenData(tokenString)

		if token.Username == "" || tokenString == "" {
			c.JSON(401, "Authentication error")
			c.Abort()
			return
		}

		c.Set("username", token.Username)
	}
}
