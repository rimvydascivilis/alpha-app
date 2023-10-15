package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(HealthCheckRoute)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code is not OK: %d", rec.Code)
	}

	expectedContentType := "application/json"
	if contentType := rec.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type header is %s; want %s", contentType, expectedContentType)
	}

	expectedMessage := "{\"message\": \"OK\"}"
	if body := rec.Body.String(); !strings.Contains(body, expectedMessage) {
		t.Errorf("Response body is %s; want %s", body, expectedMessage)
	}
}
