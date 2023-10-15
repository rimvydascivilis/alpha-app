package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	err := os.Chdir("../")
	if err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestIndexRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(indexRoute)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code is not OK: %d", rec.Code)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rec.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type header is %s; want %s", contentType, expectedContentType)
	}

	expectedMessage := "Hello, World!"
	if body := rec.Body.String(); !strings.Contains(body, expectedMessage) {
		t.Errorf("Response body is %s; want %s", body, expectedMessage)
	}
}

func TestIndexRouteWithName(t *testing.T) {
	req, err := http.NewRequest("GET", "/?name=John", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(indexRoute)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Status code is not OK: %d", rec.Code)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rec.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type header is %s; want %s", contentType, expectedContentType)
	}

	expectedMessage := "Hello, John!"
	if body := rec.Body.String(); !strings.Contains(body, expectedMessage) {
		t.Errorf("Response body is %s; want %s", body, expectedMessage)
	}
}
