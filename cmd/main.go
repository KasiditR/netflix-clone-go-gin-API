package main

import (
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/config"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/database"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true 
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	
	router.Use((gin.Logger()))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	groupRoute := router.Group("/api/v1")
	routes.MainRoutes(groupRoute)

	database.ConnectDatabase()
	log.Fatal(router.Run(":" + cfg.Port))
}
