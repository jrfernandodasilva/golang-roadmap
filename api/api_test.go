package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-roadmap/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a new recorder to capture the HTTP response
	rec := httptest.NewRecorder()

	// Create a request for the `/health` endpoint
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a handler wrapped with the `jsonMiddleware` for content-type
	handler := http.HandlerFunc(api.HealthCheckHandler)

	// Serve the request and capture the response
	handler.ServeHTTP(rec, req)

	// Check for expected status code (200 OK)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200 OK, got %d", status)
	}

	// Decode the response body
	var response api.HealthyData
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the response is `{"up": true}`
	expected := api.HealthyData{Up: true}
	if !bytes.Equal(rec.Body.Bytes(), []byte(fmt.Sprintf(`{"up":%t}`, expected.Up))) {
		t.Errorf("Expected response: %v, got: %v", expected, response) 
	}
}
