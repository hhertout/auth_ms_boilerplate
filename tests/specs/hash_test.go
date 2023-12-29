package specs

import (
	"auth_ms/src/service"
	"os"
	"testing"
)

func TestHashPassword(t *testing.T) {
	err := os.Setenv("ENCRYPT_SALT", "my_valid_salt")
	if err != nil {
		t.Error("failed to configure env")
	}
	defer os.Unsetenv("ENCRYPT_SALT")

	password := "testpassword"
	hashed, err := service.HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error hashing password: %s", err)
	}

	if hashed == "" {
		t.Error("Expected non-empty hashed password, got empty string")
	}

	err = os.Setenv("ENCRYPT_SALT", "")
	if err != nil {
		t.Error("failed to configure env")
	}
	emptySaltHashed, emptySaltErr := service.HashPassword(password)
	if emptySaltErr == nil {
		t.Error("Expected error due to empty ENCRYPT_SALT, got nil")
	}
	if emptySaltHashed != "" {
		t.Error("Expected empty hashed password, got non-empty string")
	}
}

func TestVerifyPassword(t *testing.T) {
	err := os.Setenv("ENCRYPT_SALT", "my_valid_salt")
	if err != nil {
		t.Error("failed to configure env")
	}
	defer os.Unsetenv("ENCRYPT_SALT")

	password := "testpassword"
	hashed, _ := service.HashPassword(password)
	valid, err := service.VerifyPassword(password, hashed)
	if err != nil {
		t.Errorf("Unexpected error verifying password: %s", err)
	}
	if !valid {
		t.Error("Expected valid password, got invalid")
	}

	err = os.Setenv("ENCRYPT_SALT", "")
	if err != nil {
		t.Error("failed to configure env")
	}
	emptySaltValid, emptySaltErr := service.VerifyPassword(password, hashed)
	if emptySaltErr == nil {
		t.Error("Expected error due to empty ENCRYPT_SALT, got nil")
	}
	if emptySaltValid {
		t.Error("Expected invalid password due to empty ENCRYPT_SALT, got valid")
	}

	invalidHash := "invalidhash"
	invalidHashValid, invalidHashErr := service.VerifyPassword(password, invalidHash)
	if invalidHashErr == nil {
		t.Error("Expected error due to invalid hash, got nil")
	}
	if invalidHashValid {
		t.Error("Expected invalid password due to invalid hash, got valid")
	}
}
