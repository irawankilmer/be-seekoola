package routes

import (
	"be-sakoola/handlers"
	"be-sakoola/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	auth := api.Group("/dashboard")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", handlers.Profile)

		admin := auth.Use(middleware.RoleMiddleware())
		{
			admin.GET("/post", handlers.GetPost)
		}
	}
}
