package jwt

import (
	"crypto"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
)

type PublicKeyDeterminationFunc func(keys map[string]crypto.PublicKey) string

type Validator struct {
	method jwt.SigningMethod
	keys   map[string]*rsa.PublicKey
}

func (v *Validator) Verify(raw string) (Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(raw, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrInvalidKey
		}

		pubKey, ok := v.keys[claims.KID]
		if !ok {
			return nil, jwt.ErrInvalidKey
		}

		return pubKey, nil
	})
	if err != nil {
		return Claims{}, err
	}

	return claims, err
}
