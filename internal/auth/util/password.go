package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword gets the hash of the password
func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}
	return string(hash), nil
}

// CheckPassword checks if the hashed password is correct
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
