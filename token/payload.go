package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	ExpiryAt  time.Time
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	ID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        ID,
		Username:  username,
		CreatedAt: time.Now(),
		ExpiryAt:  time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() (*Payload, error) {
	if payload.ExpiryAt.After(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return payload, nil
}
