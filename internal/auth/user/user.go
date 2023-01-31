package user

import (
	"github.com/google/uuid"
	"github.com/umalmyha/auth-server/internal/auth"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

func New(dto *NewUserDto) (*User, error) {
	// TODO: add validation, etc
	passwordHash, err := auth.GenerateHash(dto.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           uuid.NewString(),
		Email:        dto.Email,
		PasswordHash: passwordHash,
	}, nil
}
