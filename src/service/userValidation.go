package service

import (
	"errors"
)

func ValidateUserCreationData(email string, password string) error {
	if email == "" {
		return errors.New("email missing")
	}
	if password == "" {
		return errors.New("email missing")
	}

	return nil
}

func ValidatePassword(password string) error {
	// TODO
	return nil
}
