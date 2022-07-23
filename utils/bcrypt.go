package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashBytes, err := bcrypt.GenerateFromPassword(passwordBytes, 0)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func VerifyHash(password string, hashedPassword string) (bool, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	if err != nil {
		return false, err
	}
	return true, nil
}