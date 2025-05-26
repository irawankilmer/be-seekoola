package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:100;not null"`
	Body      string         `json:"body" gorm:"type:text;not null"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"user"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
