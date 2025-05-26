package main

import (
	"be-sakoola/config"
	"be-sakoola/migrations"
	"be-sakoola/models"
	"be-sakoola/routes"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.InitDB()
	migrations.RunMigrations()
	role := models.Role{Name: "admin"}
	config.DB.FirstOrCreate(&role, role)

	hPass, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	user := models.User{
		Name:     "Biasa Admin",
		Email:    "admin@example.com",
		Password: string(hPass),
		Roles:    []models.Role{role},
	}
	config.DB.FirstOrCreate(&user, models.User{Email: user.Email})

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run()
}
