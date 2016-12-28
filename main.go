package main

import (
	"log"

	"github.com/drgarcia1986/shortener/server"
	"github.com/drgarcia1986/shortener/storage"
)

func main() {
	storage := storage.NewFake()
	server := server.New(storage, "http://short.io/", 5)

	port := ":8000"
	log.Printf("Start Server at port %s", port)
	log.Fatal(server.Run(port))
}
