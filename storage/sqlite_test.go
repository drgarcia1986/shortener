package storage

import (
	"database/sql"
	"os"
	"testing"

	"github.com/drgarcia1986/shortener/url"
)

const dbPath = "./test.db"

func createTables(db *sql.DB) error {
	sqlStmt := `
	create table urls (
		short varchar(4) not null primary key,
		original text not null,
		views integer
	)`
	_, err := db.Exec(sqlStmt)
	return err
}

func createFakeData(db *sql.DB) error {
	if err := createTables(db); err != nil {
		return err
	}
	_, err := db.Exec(`
		insert into urls (short, original, views)
		values ("abc", "http://google.com", 0)
	`)
	if err != nil {
		return err
	}
	return nil
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

	storage := SQLite{path: dbPath}
	db, err := storage.getDB()
	if err != nil {
		t.Errorf("Error on get DB: %v", err)
	}
	if err = createFakeData(db); err != nil {
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

	storage := SQLite{path: dbPath}
	db, err := storage.getDB()
	if err != nil {
		t.Errorf("Error on get DB: %v", err)
	}
	if err = createTables(db); err != nil {
		t.Errorf("Error to create tables: %v", err)
	}
	_, err = storage.Get("abc")
	if err != url.NotFound {
		t.Errorf("Expected NotFound, got: %v", err)
	}
}

func TestSqliteSet(t *testing.T) {
	defer os.Remove(dbPath)

	storage := SQLite{path: dbPath}
	db, err := storage.getDB()
	if err != nil {
		t.Errorf("Error on get DB: %v", err)
	}
	if err = createTables(db); err != nil {
		t.Errorf("Error to create tables: %v", err)
	}

	expectedShort := "cba"
	expectedOriginal := "http://golang.org"
	u := &url.Url{Short: expectedShort, Original: expectedOriginal}

	err = storage.Set(u)
	if err != nil {
		t.Errorf("Error to set new url: %v", err)
	}

	reloadedUrl, err := storage.Get(u.Short)
	if err != nil {
		t.Errorf("Error on get url: %v", err)
	}
	if u.Short != reloadedUrl.Short || u.Original != reloadedUrl.Original {
		t.Errorf("Expected %s, got %s", u, reloadedUrl)
	}
}

func TestSqliteIncViews(t *testing.T) {
	defer os.Remove(dbPath)

	storage := SQLite{path: dbPath}
	db, err := storage.getDB()
	if err != nil {
		t.Errorf("Error on get DB: %v", err)
	}
	if err = createFakeData(db); err != nil {
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
