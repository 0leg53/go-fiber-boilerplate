package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashAndPassword(password string, comparadPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(comparadPassword))
	return err
}
