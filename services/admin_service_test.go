package services

import (
	"testing"
	"github.com/Desyncq/warden/shared/database"
	"github.com/stretchr/testify/assert"
)


func TestGetAdmin(t *testing.T) {
	db, _ := database.Connect()
	admin, err := GetAdmin(db)
	assert.NoError(t, err)
	assert.NotNil(t, admin)
}
