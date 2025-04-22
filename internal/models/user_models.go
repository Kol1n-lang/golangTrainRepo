package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"not null;size:255"`
	Email    string `gorm:"unique;not null;size:255"`
	Password string `gorm:"not null"`
	Cards    []Card `gorm:"foreignKey:UserID"`
}

type Card struct {
	gorm.Model
	Number string `gorm:"not null;size:16"`
	CVV    string `gorm:"not null;size:3"`
	UserID uint   `gorm:"not null"`
}
