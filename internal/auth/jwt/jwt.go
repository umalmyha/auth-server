package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Email  string   `json:"email,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type Claims struct {
	KID string `json:"kid"`
	CustomClaims
	jwt.RegisteredClaims
}
