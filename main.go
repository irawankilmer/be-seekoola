package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/be-sakoola/config"
	"github.com/irawankilmer/be-sakoola/migrations"
	"github.com/irawankilmer/be-sakoola/models"
)

func main() {
	config.InitDB()
	migrations.RunMigrations()
	role := models.Role{Name: "admin"}
	config.DB.FirstOrCreate(&role, role)

	user := models.User{
		Name:     "Biasa Admin",
		Email:    "admin@example.com",
		Password: "1234567",
		Roles:    []models.Role{role},
	}
	config.DB.FirstOrCreate(&user, models.User{Email: user.Email})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Enterprise API",
		})
	})

	r.Run()
}
