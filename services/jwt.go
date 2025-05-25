package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("desync_secret_key")

type Claims struct {
	UserID string //`json:"user_id`
	Name string //`json:"name`
	jwt.RegisteredClaims
}

type JWTService struct {
	secretKey []byte 
}

func (j *JWTService) GenerateJWTToken(userID string, name string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(j.secretKey)
}

func (j *JWTService) ValidateJWTToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func CreateJwtService() *JWTService {
	return &JWTService{
		secretKey: []byte("desync_secret_key"),
	}
}
