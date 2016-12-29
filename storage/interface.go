package storage

import "github.com/drgarcia1986/shortener/url"

type Storage interface {
	Get(string) (*url.URL, error)
	Set(*url.URL) error
	IncViews(*url.URL) error
}
