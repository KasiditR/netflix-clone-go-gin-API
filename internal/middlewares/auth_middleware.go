package middlewares

import (
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "No Cookie"})
			return
		}

		claims, msg := tokens.ValidateToken(string(cookie))
		if msg != "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)
		c.Set("image", claims.Image)
		c.Next()
	}
}
