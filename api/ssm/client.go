package ssm

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	MaxRetries = 3
	RetryDelay = 2 * time.Second
)

func DoPostWithRetry(endpoint string, form url.Values) (string, error) {
	var body string
	var err error
	for i := 1; i <= MaxRetries; i++ {
		log.Printf("POST to %s (attempt %d)", endpoint, i)
		body, err = doPost(endpoint, form)
		if err == nil {
			return body, nil
		}
		log.Printf("Attempt %d failed: %v", i, err)
		time.Sleep(RetryDelay)
	}
	return "", err
}

func doPost(endpoint string, form url.Values) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.PostForm(endpoint, form)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
