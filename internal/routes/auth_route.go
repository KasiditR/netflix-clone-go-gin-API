package routes

import (
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/handlers"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func authRoutes(routes *gin.RouterGroup) {
	routes.POST("/signup", handlers.SignUp())
	routes.POST("/login", handlers.Login())
	routes.POST("/logout", handlers.Logout())
	routes.GET("/authCheck", middlewares.Authentication(), handlers.AuthCheck())
	routes.POST("/refresh", middlewares.Authentication(), handlers.RefreshToken())
}
