package storage

import (
	"database/sql"

	"github.com/drgarcia1986/shortener/url"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	path string
}

func (s *SQLite) getDB() (*sql.DB, error) {
	return sql.Open("sqlite3", s.path)
}

func (s *SQLite) Get(short string) (*url.URL, error) {
	db, err := s.getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select original, views from urls where short = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var original string
	var views int
	err = stmt.QueryRow(short).Scan(&original, &views)
	if err != nil {
		return nil, url.ErrNotFound
	}

	return &url.URL{Short: short, Original: original, Views: views}, nil
}

func (s *SQLite) Set(u *url.URL) error {
	db, err := s.getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(
		"insert into urls (short, original, views) values (?, ?, 0)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Short, u.Original)
	return err
}

func (s *SQLite) IncViews(u *url.URL) error {
	db, err := s.getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(
		"update urls set views=views+1 where short=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Short)
	return err
}

func (s *SQLite) Create() error {
	db, err := s.getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := `
	create table urls (
		short varchar(4) not null primary key,
		original text not null,
		views integer
	)`
	_, err = db.Exec(sqlStmt)
	return err
}

func NewSQLite(path string) Storage {
	return &SQLite{path: path}
}
