package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const DefaultTimeToLive = 10 * time.Minute

type issuerConfigFunc func(cfg *issuerConfig)

type issuerConfig struct {
	issuer     string
	timeToLive time.Duration
	kidSelFn KIDSelectionFunc
}

type KIDSelectionFunc func(keys map[string]*rsa.PrivateKey) (kid string)

type KeyPair struct {
	KID        string
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

type SignClaims struct {
	Subject  string
	Audience []string
	CustomClaims
}

func WithIssuerClaim(iss string) issuerConfigFunc {
	return func(cfg *issuerConfig) {
		cfg.issuer = iss
	}
}

func WithTimeToLive(ttl time.Duration) issuerConfigFunc {
	return func(cfg *issuerConfig) {
		cfg.timeToLive = ttl
	}
}

func WithSigningKeySelectionFunc(kidSelFn KIDSelectionFunc) issuerConfigFunc {
	return func(cfg *issuerConfig) {
		if kidSelFn != nil {
			cfg.kidSelFn = kidSelFn
		}
	}
}

type Issuer struct {
	issuer     string
	method     jwt.SigningMethod
	timeToLive time.Duration
	keys       []KeyPair
	kidSelFn func() string
}

func NewIssuer(keyPairs []KeyPair, configs ...issuerConfigFunc) (*Issuer, error) {
	if len(keyPairs) == 0 {
		return nil, fmt.Errorf("there must be at least one private key provided")
	}

	cfg := &issuerConfig{
		timeToLive: DefaultTimeToLive,
		kidSelFn: defaultKIDSelectionFunc,
	}

	for _, cfgFn := range configs {
		cfgFn(cfg)
	}

	privateKeys := make(map[string]*rsa.PrivateKey)
	for _, pair := range keyPairs {
		privateKeys[pair.KID] = pair.PrivateKey
	}

	kidSelFn := func() string {
		return cfg.kidSelFn(privateKeys)
	}

	return &Issuer{
		issuer:     cfg.issuer,
		method:     jwt.SigningMethodRS256,
		timeToLive: cfg.timeToLive,
		keys:       keyPairs,
		kidSelFn: kidSelFn,
	}, nil
}

func (iss *Issuer) Sign(sc SignClaims, issueAt time.Time) error {
	kid := iss.kidSelFn()
	iss.

	expiresAt := issueAt.Add(iss.timeToLive)

	claims := &Claims{
		CustomClaims: CustomClaims{
			KID:
			Email:  sc.Email,
			Scopes: sc.Scopes,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Issuer:    iss.issuer,
			Subject:   sc.Subject,
			Audience:  sc.Audience,
			IssuedAt:  jwt.NewNumericDate(issueAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(iss.method, claims)

	token.SignedString()
}

func defaultKIDSelectionFunc(keys map[string]*rsa.PrivateKey) (kid string) {
	for kid = range keys {
		break
	}
	return kid
}