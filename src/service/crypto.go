package service

import (
	"crypto/subtle"
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

func VerifyPassword(password string, hash string) (bool, error) {
	salt := os.Getenv("ENCRYPT_SALT")
	if salt == "" {
		return false, errors.New("env variable ENCRYPT_SALT is not set")
	}

	decodeHash, err := base64.RawStdEncoding.DecodeString(hash)
	if err != nil {
		return false, err
	}

	hashToCompare := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	res := subtle.ConstantTimeCompare(hashToCompare, decodeHash)

	return res == 1, nil
}
