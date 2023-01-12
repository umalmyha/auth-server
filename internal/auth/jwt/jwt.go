package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	KID    string   `json:"kid"`
	Email  string   `json:"email,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type Claims struct {
	CustomClaims
	jwt.RegisteredClaims
}
