package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var ErrExpireToken = errors.New("token has expired")

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		UserName:  username,
		ExpiredAt: time.Now().Add(duration),
		IssuedAt:  time.Now(),
	}

	return payload, nil
}

func (payload *Payload) isValid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpireToken
	}

	return nil
}
