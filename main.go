package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drgarcia1986/shortener/server"
	"github.com/drgarcia1986/shortener/storage"
)

var toCreateStorage bool
var storageType int
var storagePath string

func init() {
	flag.BoolVar(&toCreateStorage, "create-storage", false, "To create storage and exit")
	flag.IntVar(
		&storageType,
		"storage-type",
		storage.SQLiteType,
		fmt.Sprintf("%d = Fake | %d = SQLite3", storage.FakeType, storage.SQLiteType),
	)
	flag.StringVar(&storagePath, "storage-path", "", "Path of storage (need for SQLite)")

	flag.Parse()
}

func createStorage(s storage.Storage) {
	log.Printf("Create storage on %s", storagePath)
	if err := s.Create(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("Success")
	os.Exit(0)
}

func main() {
	db := storage.New(storageType, storagePath)
	if toCreateStorage {
		createStorage(db)
	}

	s := server.New(db, "http://short.io/", 5)
	port := ":8000"
	log.Printf("Start Server at port %s", port)
	log.Fatal(s.Run(port))
}
