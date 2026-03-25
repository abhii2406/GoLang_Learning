package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPassword), err
}

func MathchPassword(pass, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	return err == nil
}
