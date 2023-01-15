package session

import "time"

type RefreshToken struct {
	ID          string
	UserID      string
	Fingerprint string
	IssuedAt    time.Time
	ExpiresAt   time.Time
}
