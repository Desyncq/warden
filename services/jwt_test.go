package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateJWT(t *testing.T) {
	user := "testuser"

	token, err := GenerateJWT(user)

	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, token)

	claims, err := ValidateJWT(token)
	if err != nil {
		t.Fatalf("Token validation failed: %v", err)
	}

	if claims.UserID != user {
		t.Errorf("Expected userID %s, got %s", user, claims.UserID)
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		t.Error("Token already expired")
	}

	assert.NoError(t, err)
}

func TestInvalidJWT(t *testing.T) {
	_, err := ValidateJWT("invalid.token.string")
	if err == nil {
		t.Error("Expected error for invalid token, got nil")
	}
}
