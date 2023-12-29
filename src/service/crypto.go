package service

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
	"os"
	"time"
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

func GenerateJwtToken(user string) (string, error) {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return "", errors.New("env variable JWT_SECRET is not set")
	}

	expiresAt := time.Now().Unix() + 3600*24*20

	claims := jwt.RegisteredClaims{
		Issuer:    user,
		ExpiresAt: jwt.NewNumericDate(time.Unix(expiresAt, 0)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJwtToken(tokenString string) (bool, error) {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		return false, errors.New("env variable JWT_SECRET is not set")
	}

	var claims jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false, err
	}

	valid := token.Valid

	return valid, nil
}
