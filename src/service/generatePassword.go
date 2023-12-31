package service

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomPassword() (string, error) {
	passwordLength := 16
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*=+?/"
	charsetLen := big.NewInt(int64(len(charset)))

	password := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}
