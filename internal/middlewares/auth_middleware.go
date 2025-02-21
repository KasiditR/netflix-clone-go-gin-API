package middlewares

import (
	"net/http"

	"github.com/KasiditR/netflix-clone-go-gin-API/internal/tokens"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "No Authorization Header"})
			c.Abort()
			return
		}

		claims, err := tokens.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			c.Abort()
			return
		}

		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)
		c.Set("image", claims.Image)
		c.Next()
	}
}
