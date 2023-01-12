package jwt

import (
	"crypto"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type PublicKeyDeterminationFunc func(keys map[string]crypto.PublicKey) string

type Validator struct {
	method jwt.SigningMethod
	keys map[string]crypto.PublicKey
}

func (v *Validator) Verify(raw string) (Claims, error) {
	time.Time{}
	var claims Claims
	if _, err := jt.
}