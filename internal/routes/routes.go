package routes

import (
	"github.com/gin-gonic/gin"
)

func MainRoutes(routes *gin.RouterGroup) {
	authRoute := routes.Group("/auth")
	authRoutes(authRoute)

	contentRoute := routes.Group("/content")
	contentRoutes(contentRoute)

	searchRoute := routes.Group("/search")
	searchRoutes(searchRoute)
}
