package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const DefaultTimeToLive = 10 * time.Minute

var DefaultKeyDeterminationFunc = func(pairs []KeyPair) (kp KeyPair) {
	return pairs[0]
}

type KeyDeterminationFunc func([]KeyPair) KeyPair

type issuerConfigFunc func(cfg *issuerConfig)

type issuerConfig struct {
	issuer     string
	timeToLive time.Duration
	keyFunc    KeyDeterminationFunc
}

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

func WithSigningKeySelectionFunc(keyFunc KeyDeterminationFunc) issuerConfigFunc {
	return func(cfg *issuerConfig) {
		if keyFunc != nil {
			cfg.keyFunc = keyFunc
		}
	}
}

type Issuer struct {
	issuer     string
	method     jwt.SigningMethod
	timeToLive time.Duration
	keyPairs   []KeyPair
	keyFunc    KeyDeterminationFunc
}

func NewIssuer(keyPairs []KeyPair, configs ...issuerConfigFunc) (*Issuer, error) {
	if len(keyPairs) == 0 {
		return nil, fmt.Errorf("there must be at least one key pair provided")
	}

	cfg := &issuerConfig{
		timeToLive: DefaultTimeToLive,
		keyFunc:    DefaultKeyDeterminationFunc,
	}

	for _, cfgFn := range configs {
		cfgFn(cfg)
	}

	return &Issuer{
		issuer:     cfg.issuer,
		method:     jwt.SigningMethodRS256,
		timeToLive: cfg.timeToLive,
		keyPairs:   keyPairs,
		keyFunc:    cfg.keyFunc,
	}, nil
}

func (iss *Issuer) Sign(sc SignClaims, issueAt time.Time) (string, error) {
	kp := iss.keyFunc(iss.keyPairs)
	if kp.KID == "" {
		return "", errors.New("failed to determine kid with provided key determination func")
	}

	expiresAt := issueAt.Add(iss.timeToLive)

	claims := Claims{
		KID: kp.KID,
		CustomClaims: CustomClaims{
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

	token, err := jwt.NewWithClaims(iss.method, claims).SignedString(kp.PrivateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (iss *Issuer) Validator() *Validator {
	keys := make(map[string]*rsa.PublicKey)
	for _, kp := range iss.keyPairs {
		keys[kp.KID] = kp.PublicKey
	}

	return &Validator{
		method: iss.method,
		keys:   keys,
	}
}
