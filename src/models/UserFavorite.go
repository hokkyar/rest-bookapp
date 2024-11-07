package models

import (
	"time"

	"gorm.io/gorm"
)

type UserFavorite struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	BookID    uint           `gorm:"not null" json:"book_id"`
	Book      Book           `gorm:"foreignKey:BookID" json:"book"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
