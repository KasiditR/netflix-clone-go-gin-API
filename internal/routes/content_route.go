package routes

import (
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/handlers"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/middlewares"
	"github.com/gin-gonic/gin"
	"slices"
)

func contentRoutes(routes *gin.RouterGroup) {
	routes.Use(middlewares.Authentication())
	routes.Use(func(c *gin.Context) {
		contentType := c.Query("contentType")
		types := []string{"movie", "tv"}
		if contentType == "" || !slices.Contains(types, contentType) {
			c.AbortWithStatusJSON(400, gin.H{"message": "contentType not found. contentType must contain (movie, tv)"})
			return
		}
		c.Next()
	})
	routes.GET("/trending", handlers.GetContentTrading())
	routes.GET("/trailers", handlers.GetContentTrailers())
	routes.GET("/detail", handlers.GetContentDetail())
	routes.GET("/similar", handlers.GetContentSimilar())
	routes.GET("/category", handlers.GetContentByCategory())
}
