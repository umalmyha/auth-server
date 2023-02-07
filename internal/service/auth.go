package service

import (
	"context"
	"time"

	"github.com/umalmyha/auth-server/internal/auth"
	"github.com/umalmyha/auth-server/internal/auth/jwt"
	"github.com/umalmyha/auth-server/internal/auth/session"
	"github.com/umalmyha/auth-server/internal/auth/user"
)

// TODO: error handling + wrapping

type UserRepository interface {
	CreateUser(ctx context.Context, user *user.User) error
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}

type RefreshTokenRepository interface {
	GetAllByUserID(ctx context.Context, userID string) ([]*session.RefreshToken, error)
}

type AuthService struct {
	issuer           *jwt.Issuer
	userRepo         UserRepository
	refreshTokenRepo RefreshTokenRepository
}

func NewAuthService(
	issuer *jwt.Issuer,
	userRepo UserRepository,
	refreshTokenRepo RefreshTokenRepository,
) *AuthService {
	return &AuthService{
		issuer:           issuer,
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *AuthService) Register(ctx context.Context, nu *user.NewUserDto) error {
	user, err := user.New(nu)
	if err != nil {
		return err
	}

	if err = s.userRepo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, login *auth.LoginDto, now time.Time) error {
	user, err := s.userRepo.GetByEmail(ctx, login.Email)
	if err != nil {
		return err
	}

	if err = auth.VerifyPassword(user.PasswordHash, login.Password); err != nil {
		return err
	}

	tokens, err := s.refreshTokenRepo.GetAllByUserID(ctx, user.ID)
	if err != nil {
		return err
	}

	claims := jwt.SignClaims{
		Subject: user.ID,
		CustomClaims: jwt.CustomClaims{
			Email:  user.Email,
			Scopes: ,
		},
	}

	accessToken, err := s.issuer.Sign(claims, now)
	if err != nil {
		return err
	}
}
