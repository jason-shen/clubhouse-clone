package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ComparePassword(attempt string, password string) error {
	bytePassword, byteHashedPasword := []byte(attempt), []byte(password)
	return bcrypt.CompareHashAndPassword(byteHashedPasword, bytePassword)
}
