package migrations

import (
	"be-sakoola/internal/config"
	"be-sakoola/models"
	"log"
)

func RunMigrations() {
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Post{},
	)

	if err != nil {
		log.Fatal("Migrasi gagal!")
	}

	log.Println("Migrasi Database berhasil!")
}
