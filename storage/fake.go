package storage

import "github.com/drgarcia1986/shortener/url"

type Fake struct {
	KnowUrls map[string]*url.URL
}

func (f *Fake) Get(short string) (*url.URL, error) {
	u, found := f.KnowUrls[short]
	if !found {
		return nil, url.ErrNotFound
	}
	return u, nil
}

func (f *Fake) Set(u *url.URL) error {
	f.KnowUrls[u.Short] = u
	return nil
}

func (f *Fake) IncViews(u *url.URL) error {
	u, found := f.KnowUrls[u.Short]
	if found {
		u.Views++
	}
	return nil
}

func NewFake() Storage {
	urls := make(map[string]*url.URL)
	return &Fake{KnowUrls: urls}
}
