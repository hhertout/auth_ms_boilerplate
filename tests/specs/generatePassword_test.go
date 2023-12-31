package specs

import (
	"auth_ms/src/service"
	"strings"
	"testing"
)

func TestGenerateRandomPassword(t *testing.T) {
	password, err := service.GenerateRandomPassword()
	if err != nil {
		t.Errorf("Failed to generate password : %s", err)
	}

	expectedLength := 16
	if len(password) != expectedLength {
		t.Errorf("Incorrect password length : %d, receive : %d", expectedLength, len(password))
	}

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*=+?/"
	for _, char := range password {
		if !strings.ContainsRune(charset, char) {
			t.Errorf("Unauthorized char found in random password gen : %s", string(char))
		}
	}
}
