package storage

import "github.com/drgarcia1986/shortener/url"

type Fake struct {
	KnowUrls map[string]*url.Url
}

func (f *Fake) Get(short string) (*url.Url, error) {
	u, found := f.KnowUrls[short]
	if !found {
		return nil, url.NotFound
	}
	return u, nil
}

func (f *Fake) Set(u *url.Url) error {
	f.KnowUrls[u.Short] = u
	return nil
}

func (f *Fake) IncViews(u *url.Url) error {
	u, found := f.KnowUrls[u.Short]
	if found {
		u.Views++
	}
	return nil
}

func NewFake() Storage {
	urls := make(map[string]*url.Url)
	return &Fake{KnowUrls: urls}
}
