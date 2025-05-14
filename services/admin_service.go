package services

import (
	"gorm.io/gorm"
	"github.com/Desyncq/warden/models"
	"log"
)

func GetAdmin(db *gorm.DB) (models.Admin, error) {
	var Admins []models.Admin
	result := db.Raw("SELECT * FROM Admin").Scan(&Admins)
	if result.Error != nil {
		log.Fatal("Failed to get admins: ", result.Error)
	}
	return Admins[0], result.Error
}

