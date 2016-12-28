package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drgarcia1986/shortener/url"
)

func TestStatsHandlerSuccessful(t *testing.T) {
	server, fakeStorage := getServerAndStorage()

	short := "abc"
	expectedViews := 2
	err := fakeStorage.Set(&url.Url{Short: short, Original: "http://golang.org", Views: expectedViews})
	if err != nil {
		t.Errorf("Error on create a url: %v", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/s/%s/", short), nil)
	if err != nil {
		t.Errorf("Error on create a fake request: %v", err)
	}

	w := httptest.NewRecorder()
	server.router().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, w.Code)
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("Error on read recorded response body: %v", err)
	}
	var u url.Url
	err = json.Unmarshal(body, &u)
	if err != nil {
		t.Errorf("Error to unmarshal response body: %v", err)
	}

	if u.Views != expectedViews {
		t.Errorf("Expected %d, got %d", expectedViews, u.Views)
	}
}

func TestStatsHandlerNotFound(t *testing.T) {
	server, _ := getServerAndStorage()

	req, err := http.NewRequest("GET", "/s/abc/", nil)
	if err != nil {
		t.Errorf("Error on create a fake request: %v", err)
	}

	w := httptest.NewRecorder()
	server.router().ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected %d, got %d", http.StatusNotFound, w.Code)
	}
}
