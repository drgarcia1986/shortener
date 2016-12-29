package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/drgarcia1986/shortener/url"
)

func TestRedirectHandlerSuccessful(t *testing.T) {
	server, fakeStorage := getServerAndStorage()

	short := "abc"
	expectedOriginal := "http://golang.org"
	err := fakeStorage.Set(&url.URL{Short: short, Original: expectedOriginal})
	if err != nil {
		t.Errorf("Error on create a url: %v", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/", short), nil)
	if err != nil {
		t.Errorf("Error on create a fake request: %v", err)
	}

	w := httptest.NewRecorder()
	server.router().ServeHTTP(w, req)

	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected %d, got %d", http.StatusMovedPermanently, w.Code)
	}

	location := w.Header().Get("Location")
	if location != expectedOriginal {
		t.Errorf("Expected %s, got %s", expectedOriginal, location)
	}
}

func TestRedirectHandlerIncUrlViews(t *testing.T) {
	server, fakeStorage := getServerAndStorage()

	short := "abc"
	currentViews := 1
	expectedOriginal := "http://golang.org"
	err := fakeStorage.Set(&url.URL{Short: short, Original: expectedOriginal, Views: currentViews})
	if err != nil {
		t.Errorf("Error on create a url: %v", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s/", short), nil)
	if err != nil {
		t.Errorf("Error on create a fake request: %v", err)
	}

	w := httptest.NewRecorder()
	server.router().ServeHTTP(w, req)

	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected %d, got %d", http.StatusMovedPermanently, w.Code)
	}

	time.Sleep(5 * time.Millisecond)
	u, err := fakeStorage.Get(short)
	if err != nil {
		t.Errorf("Error to get url: %v", err)
	}

	expectedViews := currentViews + 1
	if u.Views != expectedViews {
		t.Errorf("Expected %d, got %d", expectedViews, u.Views)
	}
}

func TestRedirectHandlerNotFound(t *testing.T) {
	server, _ := getServerAndStorage()

	req, err := http.NewRequest("GET", "/abc/", nil)
	if err != nil {
		t.Errorf("Error on create a fake request: %v", err)
	}

	w := httptest.NewRecorder()
	server.router().ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected %d, got %d", http.StatusNotFound, w.Code)
	}
}
