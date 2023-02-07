package session

import "time"

const DefaultTimeToLive = 24 * 30 * time.Hour // 30 days

type originatorConfigFunc func()

type Originator struct {
	maxTokensCount int
	timeToLive     time.Duration
}

func NewOriginator() *Originator {
	return &Originator{
		maxTokensCount: 0,
		timeToLive:     0,
	}
}
