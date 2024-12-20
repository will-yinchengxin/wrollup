package es

import (
	"fmt"
	"net/http"
	"time"
)

func CheckESConnection(url string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to connect to Elasticsearch: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("Elasticsearch returned unexpected status code: %d", resp.StatusCode)
}
