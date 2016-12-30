package storage

import (
	"os"
	"testing"

	"github.com/drgarcia1986/shortener/url"
)

const dbPath = "./test.db"

func createFakeData(s *SQLite) error {
	if err := s.Create(); err != nil {
		return err
	}
	db, err := s.getDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		insert into urls (short, original, views)
		values ("abc", "http://google.com", 0)
	`)
	if err != nil {
		return err
	}
	return nil
}

func TestSqliteCreate(t *testing.T) {
	storage := &SQLite{path: dbPath}
	if err := storage.Create(); err != nil {
		t.Errorf("Error on create storage: %v", err)
	}
}

func TestSqliteGetDB(t *testing.T) {
	defer os.Remove(dbPath)

	storage := SQLite{path: dbPath}
	if _, err := storage.getDB(); err != nil {
		t.Errorf("Error on get DB: %v", err)
	}
}

func TestSqliteGet(t *testing.T) {
	defer os.Remove(dbPath)

	storage := &SQLite{path: dbPath}
	if err := createFakeData(storage); err != nil {
		t.Errorf("Error on create fake data: %v", err)
	}
	u, err := storage.Get("abc")
	if err != nil {
		t.Errorf("Error on get url: %v", err)
	}

	expectedShort := "abc"
	expectedOriginal := "http://google.com"
	if u.Short != expectedShort {
		t.Errorf("Expected %s, got %s", expectedShort, u.Short)
	}
	if u.Original != expectedOriginal {
		t.Errorf("Expected %s, got %s", expectedOriginal, u.Original)
	}
}

func TestSqliteGetNotFound(t *testing.T) {
	defer os.Remove(dbPath)

	storage := &SQLite{path: dbPath}
	if err := storage.Create(); err != nil {
		t.Errorf("Error to create tables: %v", err)
	}
	_, err := storage.Get("abc")
	if err != url.ErrNotFound {
		t.Errorf("Expected NotFound, got: %v", err)
	}
}

func TestSqliteSet(t *testing.T) {
	defer os.Remove(dbPath)

	storage := &SQLite{path: dbPath}
	if err := storage.Create(); err != nil {
		t.Errorf("Error to create tables: %v", err)
	}

	expectedShort := "cba"
	expectedOriginal := "http://golang.org"
	u := &url.URL{Short: expectedShort, Original: expectedOriginal}

	if err := storage.Set(u); err != nil {
		t.Errorf("Error to set new url: %v", err)
	}

	reloadedURL, err := storage.Get(u.Short)
	if err != nil {
		t.Errorf("Error on get url: %v", err)
	}
	if u.Short != reloadedURL.Short || u.Original != reloadedURL.Original {
		t.Errorf("Expected %s, got %s", u, reloadedURL)
	}
}

func TestSqliteIncViews(t *testing.T) {
	defer os.Remove(dbPath)

	storage := &SQLite{path: dbPath}
	if err := createFakeData(storage); err != nil {
		t.Errorf("Error on create fake data: %v", err)
	}
	u, err := storage.Get("abc")
	if err != nil {
		t.Errorf("Error on get url: %v", err)
	}

	oldViewCount := u.Views
	if err = storage.IncViews(u); err != nil {
		t.Errorf("Error on IncViews: %v", err)
	}
	u, err = storage.Get("abc")
	if err != nil {
		t.Errorf("Error on get url: %v", err)
	}

	if u.Views <= oldViewCount {
		t.Errorf("Expected %d, got %d", oldViewCount+1, u.Views)
	}
}
