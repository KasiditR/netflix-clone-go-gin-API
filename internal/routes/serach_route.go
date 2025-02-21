package routes

import (
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/handlers"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func searchRoutes(routes *gin.RouterGroup) {
	routes.Use(middlewares.Authentication())
	routes.GET("", handlers.SearchContent())
	routes.GET("/history", handlers.GetSearchHistory())
	routes.DELETE("/remove/:id", handlers.RemoveItemFromSearchHistory())
	routes.DELETE("/clear", handlers.ClearSearchHistory())
}
