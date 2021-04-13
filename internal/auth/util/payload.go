package util

import (
	"errors"
	"time"
)

// Error returned by token
var (
	ErrExpiredToken = errors.New("Token has expired")
	ErrInvalidToken = errors.New("Token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with email and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		Email:     username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the payload is valid
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
