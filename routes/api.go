package routes

import (
	"be-sakoola/handlers"
	"be-sakoola/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	auth := api.Group("/dashboard")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", handlers.Profile)
		auth.GET("/post", handlers.GetAllPosts)

		admin := auth.Group("")
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			admin.POST("/post", handlers.CreatePost)
			admin.PUT("/post/:id", handlers.UpdatePost)
			admin.DELETE("/post/:id", handlers.DeletePost)
		}
	}

	// Cetak semua route yang sudah terdaftar
	// Hanya untuk development
	for _, route := range r.Routes() {
		log.Printf("%-6s -> %s", route.Method, route.Path)
	}
}
