package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyValid(t *testing.T) {
	// 1. Set up the input
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret123")

	// 2. Call the function
	key, err := GetAPIKey(headers)

	// 3. Verify the result
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if key != "secret123" {
		t.Fatalf("expected 'secret123', got: %v", key)
	}
}

func TestGetAPIKeyInvalid(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatalf("expected an err, but got nil")
	}
}
