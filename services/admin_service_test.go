package services

import (
	"github.com/Desyncq/warden/shared/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAdmin(t *testing.T) {
	db, _ := database.Connect()
	adminService := NewAdminService(db)
	admin, _ := adminService.GetAdmin()
	if admin == nil {
		assert.Nil(t, admin)
	} else {
		assert.NotNil(t, admin)
	}
}
