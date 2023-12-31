package specs

import (
	"auth_ms/src/service"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"testing"
	"time"
)

func TestGenerateJwtToken(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "my_valid_secret_key")
	if err != nil {
		t.Error("failed to configure env")
	}
	defer os.Unsetenv("JWT_SECRET")

	user := "testuser"
	token, err := service.GenerateJwtToken(user)
	if err != nil {
		t.Errorf("Unexpected error generating token: %s", err)
	}

	if token == "" {
		t.Error("Expected non-empty token, got empty string")
	}

	err = os.Setenv("JWT_SECRET", "")
	if err != nil {
		t.Error("failed to configure env")
	}
	emptyKeyToken, emptyKeyErr := service.GenerateJwtToken(user)
	if emptyKeyErr == nil {
		t.Error("Expected error due to empty JWT_SECRET, got nil")
	}
	if emptyKeyToken != "" {
		t.Error("Expected empty token, got non-empty string")
	}

	err = os.Setenv("JWT_SECRET", "my_valid_secret_key")
	if err != nil {
		t.Error("failed to configure env")
	}
	expiresAt := time.Now().Unix() + 3600*24*20
	futureExpiresToken, futureExpiresErr := service.GenerateJwtToken(user)
	if futureExpiresErr != nil {
		t.Errorf("Unexpected error generating token with future expiration: %s", futureExpiresErr)
	}

	parsedToken, parseErr := jwt.Parse(futureExpiresToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_valid_secret_key"), nil
	})
	if parseErr != nil {
		t.Errorf("Error parsing token: %s", parseErr)
	}

	exp, err := parsedToken.Claims.GetExpirationTime()
	if exp.Unix() != expiresAt {
		t.Errorf("Expected expiration time %d, got %d", expiresAt, exp.Unix())
	}
}

func TestVerifyJwtToken(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "my_valid_secret_key")
	if err != nil {
		t.Error("failed to configure env")
	}
	defer os.Unsetenv("JWT_SECRET")

	user := "testuser"
	token, _ := service.GenerateJwtToken(user)

	valid, _, err := service.VerifyJwtToken(token)
	if err != nil {
		t.Errorf("Unexpected error verifying token: %s", err)
	}

	if !valid {
		t.Error("Expected valid token, got invalid")
	}

	err = os.Setenv("JWT_SECRET", "")
	if err != nil {
		t.Error("failed to configure env")
	}
	validEmptyKey, _, errEmptyKey := service.VerifyJwtToken(token)
	if errEmptyKey == nil {
		t.Error("Expected error due to empty JWT_SECRET, got nil")
	}
	if validEmptyKey {
		t.Error("Expected invalid token due to empty JWT_SECRET, got valid")
	}

	invalidToken := "invalidtoken"
	validInvalidToken, _, errInvalidToken := service.VerifyJwtToken(invalidToken)
	if errInvalidToken == nil {
		t.Error("Expected error due to invalid token, got nil")
	}
	if validInvalidToken {
		t.Error("Expected invalid token, got valid")
	}
}
