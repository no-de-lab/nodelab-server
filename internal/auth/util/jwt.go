package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	c "github.com/no-de-lab/nodelab-server/config"
)

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(config *c.Configuration) JWTMaker {
	return JWTMaker{config.Secret.SecretKey}
}

// CreateToken creates a new token with email and duration
func (jm *JWTMaker) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(jm.secretKey))
}

// VerifyToken verifies if the token is valid
func (jm *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// Verifies signing key is valid
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jm.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
