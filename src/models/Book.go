package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Publisher   string         `gorm:"not null" json:"publisher"`
	Year        int            `gorm:"not null" json:"year"`
	Description string         `gorm:"type:text" json:"description"`
	UserID      uint           `gorm:"not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
