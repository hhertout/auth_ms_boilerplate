package service

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/argon2"
	"os"
)

func HashPassword(password string) (string, error) {
	salt := os.Getenv("ENCRYPT_SALT")
	if salt == "" {
		return "", errors.New("env variable ENCRYPT_SALT is not set")
	}
	key := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)

	return base64.RawStdEncoding.EncodeToString(key), nil
}

func VerifyPassword(password string) (bool, error) {
	return true, nil
}
