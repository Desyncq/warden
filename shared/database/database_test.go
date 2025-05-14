package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test01DBConnection(t *testing.T) {
	db, _ := Connect()
	psdb, err := db.DB()
	assert.NoError(t, err)

	err = psdb.Ping()
	assert.NoError(t, err)
}
