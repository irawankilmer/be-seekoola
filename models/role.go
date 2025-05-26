package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:50; unique; not null" json:"name"`
	Users     []User         `gorm:"many2many:user_roles;" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
