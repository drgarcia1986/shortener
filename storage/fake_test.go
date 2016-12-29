package storage

import (
	"testing"

	"github.com/drgarcia1986/shortener/url"
)

func TestFakeGet(t *testing.T) {
	short := "abc"
	expectedOriginal := "http://golang.org"

	storage := &Fake{
		KnowUrls: map[string]*url.URL{
			short: &url.URL{Short: short, Original: expectedOriginal}},
	}

	u, err := storage.Get(short)
	if err != nil {
		t.Errorf("Error on get url: %v", err)
	}

	if u.Original != expectedOriginal {
		t.Errorf("Expected %s, got %s", expectedOriginal, u.Original)
	}
}

func TestFakeGetNotFound(t *testing.T) {
	storage := &Fake{KnowUrls: map[string]*url.URL{}}

	_, err := storage.Get("abc")
	if err != url.ErrNotFound {
		t.Errorf("Expected Not Found, got: %v", err)
	}
}

func TestFakeSet(t *testing.T) {
	short := "abc"
	u := url.URL{Short: short, Original: "http://golang.org"}

	storage := &Fake{KnowUrls: map[string]*url.URL{}}
	storage.Set(&u)

	_, found := storage.KnowUrls[short]
	if !found {
		t.Error("Fake Storage don't set url")
	}
}

func TestFakeIncViews(t *testing.T) {
	short := "abc"

	u := url.URL{Short: short, Original: "http://golang.org", Views: 2}
	storage := &Fake{KnowUrls: map[string]*url.URL{short: &u}}

	err := storage.IncViews(&u)
	if err != nil {
		t.Errorf("Error on inc views: %v", err)
	}

	if u.Views != 3 {
		t.Errorf("Expected 3, got %d", u.Views)
	}
}
