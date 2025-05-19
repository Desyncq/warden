package services

import (
	"errors"
	"gorm.io/gorm"
	"github.com/Desyncq/warden/models"
)

type AdminService struct {
	db *gorm.DB
}

func (s *AdminService) AddNewAdmin(admin models.Admin) error {
	result := s.db.Create(&admin)
	return result.Error
}

func (s *AdminService) GetAdmin() (*models.Admin , error) {
	var admin models.Admin
	result := s.db.First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin not found")
		}
		return nil, result.Error
	}

	return &admin, nil
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{
		db: db,
	}
}
