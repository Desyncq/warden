package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UID uint `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

