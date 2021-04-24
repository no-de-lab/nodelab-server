package util

import (
	"testing"
	"time"

	"github.com/no-de-lab/nodelab-server/config"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	secretInfo := config.SecretInfo{SecretKey: "1234567890123456789123456789012"}
	config := config.Configuration{Secret: secretInfo}

	jm := NewJWTMaker(&config)

	username := "bob"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := jm.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := jm.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, username, payload.Email)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	secretInfo := config.SecretInfo{SecretKey: "1234567890123456789123456789012"}
	config := config.Configuration{Secret: secretInfo}

	jm := NewJWTMaker(&config)

	token, err := jm.CreateToken("bob", -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := jm.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
