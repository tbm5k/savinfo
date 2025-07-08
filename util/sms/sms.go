package sms

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Sms struct {
	username string
	apiKey string
	apiURL string
}

func NewSmser() *Sms {
	return &Sms{
		username: "sandbox",
		apiKey: "MyAppApiKey",
		apiURL: "https://api.africastalking.com/version1/messaging/bulk",
	}
}

func (s *Sms) Send(m, r string) error {

	p := fmt.Sprintf(`{
		"username": "%s",
		"message": "%s",
		"senderId": "%s",
		"phoneNumbers": ["%s"]
	}`, s.username, m, "abc", r)

	req, err := http.NewRequest("POST", s.apiURL, bytes.NewBuffer([]byte(p)))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apiKey", s.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return err
	}

	log.Println(b)

	return nil
}

