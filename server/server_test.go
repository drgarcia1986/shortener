package server

import "github.com/drgarcia1986/shortener/storage"

func getServerAndStorage() (*Server, *storage.Fake) {
	fakeStorage := storage.NewFake()
	return New(fakeStorage, "http://test.io", 5), fakeStorage.(*storage.Fake)
}
