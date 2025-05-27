package routes

import (
	handlers2 "be-sakoola/internal/handlers"
	middleware2 "be-sakoola/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", handlers2.Register)
	api.POST("/login", handlers2.Login)

	auth := api.Group("/dashboard")
	auth.Use(middleware2.AuthMiddleware())
	{
		auth.GET("/profile", handlers2.Profile)
		auth.GET("/post", handlers2.GetAllPosts)

		admin := auth.Group("")
		admin.Use(middleware2.RoleMiddleware("admin"))
		{
			admin.POST("/post", handlers2.CreatePost)
			admin.PUT("/post/:id", handlers2.UpdatePost)
			admin.DELETE("/post/:id", handlers2.DeletePost)
		}
	}

	// Cetak semua route yang sudah terdaftar
	// Hanya untuk development
	for _, route := range r.Routes() {
		log.Printf("%-6s -> %s", route.Method, route.Path)
	}
}
