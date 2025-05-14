package services

import (
	"errors"
	"gorm.io/gorm"
	"github.com/Desyncq/warden/models"
	"log"
)

func GetAdmin(db *gorm.DB) (models.Admin, error) {
	var Admins []models.Admin
	result := db.Find(&Admins)
	if result.Error != nil {
		log.Fatal("Failed to get admins: ", result.Error)
	}

	if len(Admins) == 0 {
		return models.Admin{}, errors.New("no admins found")
	}

	return Admins[0], result.Error
}

