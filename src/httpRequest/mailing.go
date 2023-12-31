package httpRequest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

func SendReinitialisationMail(to string, password string) error {
	mailerUrl := os.Getenv("MAILER_URL")
	if mailerUrl == "" {
		return errors.New("mailer url is not set")
	}

	type RequestBody struct {
		Password string `json:"password"`
	}

	type Request struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    RequestBody
	}

	request := Request{
		To:      to,
		Subject: "Password Reinitialisation",
		Body: RequestBody{
			Password: password,
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", mailerUrl+"/api/mailer/password-reset", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(r)
	if err != nil {
		return err
	}

	return nil
}

func SendPasswordUpdatedMail(to string) error {
	mailerUrl := os.Getenv("MAILER_URL")
	if mailerUrl == "" {
		return errors.New("mailer url is not set")
	}

	type RequestBody struct{}

	type Request struct {
		To      []string `json:"to"`
		Subject string   `json:"subject"`
	}

	request := Request{
		To:      []string{to},
		Subject: "Your password have been updated",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", mailerUrl+"/api/mailer/password-updated", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(r)
	if err != nil {
		return err
	}

	return nil
}
