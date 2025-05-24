package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPaswordAndCheckPasswordHash(t *testing.T) {
	pass := "testpass"

	hash, err := HashPassword(pass)
	if err != nil {
		t.Fatalf("Couldn't hash password: %v", err)
	}

	val := CheckPasswordHash(pass, hash)
	if val != true {
		t.Error("Password is incorrect")
	}

	assert.NoError(t, err)
}
