package storage

import (
	"sync"

	"github.com/drgarcia1986/shortener/url"
)

type Fake struct {
	KnowUrls map[string]*url.URL
	mutex    *sync.RWMutex
}

func (f *Fake) Get(short string) (*url.URL, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
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
	f.mutex.Lock()
	defer f.mutex.Unlock()
	u.Views++
	return nil
}

func (s *Fake) Create() error {
	return nil
}

func NewFake() Storage {
	urls := make(map[string]*url.URL)
	mutex := &sync.RWMutex{}
	return &Fake{KnowUrls: urls, mutex: mutex}
}
