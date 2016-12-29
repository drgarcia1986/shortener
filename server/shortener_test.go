package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/drgarcia1986/shortener/url"
)

func TestShortenerHandlerSuccessful(t *testing.T) {
	server, fakeStorage := getServerAndStorage()

	jsonBody := `{"url": "http://golang.org"}`
	reader := strings.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/short/", reader)
	if err != nil {
		t.Errorf("Error on create a fake request: %v", err)
	}

	w := httptest.NewRecorder()
	server.shortener(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d, got %d", http.StatusCreated, w.Code)
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Error on read recorded response body: %v", err)
	}
	var u url.URL
	err = json.Unmarshal(body, &u)
	if err != nil {
		t.Errorf("Error to unmarshal response body: %v", err)
	}

	if _, ok := fakeStorage.KnowUrls[u.Short]; !ok {
		t.Errorf("Cannot persist shorted url actual urls %v", fakeStorage.KnowUrls)
	}
}

func TestShortenerHandlerBadRequest(t *testing.T) {
	server, _ := getServerAndStorage()

	jsonBody := `{"url": ""}`
	reader := strings.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/short/", reader)
	if err != nil {
		t.Errorf("Error on create a fake request: %v", err)
	}

	w := httptest.NewRecorder()
	server.shortener(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected %d, got %d", http.StatusBadRequest, w.Code)
	}
}
