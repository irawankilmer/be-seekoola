package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/be-sakoola/handler"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		v1.GET("/posts", handler.GetPosts)
		v1.POST("/posts", handler.CreatePost)
	}
}
