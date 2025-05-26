package migrations

import (
	"github.com/irawankilmer/be-sakoola/config"
	"github.com/irawankilmer/be-sakoola/models"
	"log"
)

func RunMigrations() {
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
	)

	if err != nil {
		log.Fatal("Migrasi gagal!")
	}

	log.Println("Migrasi Database berhasil!")
}
