package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}
