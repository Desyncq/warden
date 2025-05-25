package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	pass = "testpass"
	hashService = &HashService{}
)

func TestPasswordHashing(t *testing.T) {
	hash, err := hashService.HashPassword(pass)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	
}

func TestPasswordMatch(t *testing.T) {
	hash, _ := hashService.HashPassword(pass)

	match := hashService.CheckPasswordHash(pass, hash)
	assert.True(t, match)
}
