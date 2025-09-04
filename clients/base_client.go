package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type baseClient struct{}

type BaseClient interface {
	Do(httpMethod string, requestBody interface{}, baseURL string, queryParams map[string]string, headers map[string]string) (string, error)
}

func NewBaseClient() BaseClient {
	return &baseClient{}
}

func (b baseClient) Do(httpMethod string, requestBody interface{}, baseURL string, queryParams map[string]string, headers map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	var bodyBytes io.Reader
	if requestBody != nil {
		b, err := json.Marshal(requestBody)
		if err != nil {
			return "", fmt.Errorf("failed to marshal body: %w", err)
		}
		bodyBytes = strings.NewReader(string(b))
	}
	req, err := http.NewRequest(httpMethod, u.String(), bodyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// HTTP client with timeout
	client := &http.Client{Timeout: 15 * time.Second}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", errors.New(fmt.Sprintf("non-2xx status: %d, body: %s", resp.StatusCode, strings.TrimSpace(string(body))))
	}

	return string(body), nil
}
