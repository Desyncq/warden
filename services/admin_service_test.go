package services

import (
	"testing"
	"github.com/Desyncq/warden/shared/database"
	"github.com/stretchr/testify/assert"
)


func TestGetAdmin(t *testing.T) {
	db, _ := database.Connect()
	admin, _ := GetAdmin(db)
	assert.NotNil(t, admin)
}
