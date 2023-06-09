package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Claims interface {
	Valid() error
}

type Payload struct {
	UserID    uuid.UUID
	UserRole  []uuid.UUID
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errors.New("expired")
	}
	return nil
}