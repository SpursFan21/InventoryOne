package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT creates a JWT token for a user
func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
