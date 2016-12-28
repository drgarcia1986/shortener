package storage

import "github.com/drgarcia1986/shortener/url"

type Storage interface {
	Get(string) (*url.Url, error)
	Set(*url.Url) error
	IncViews(*url.Url) error
}
