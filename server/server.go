package server

import (
	"net/http"

	"github.com/drgarcia1986/shortener/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	db          storage.Storage
	urlBase     string
	shortLength int
}

func (s *Server) router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/short/", s.shortener).Methods("POST")
	r.HandleFunc("/s/{short}/", s.stats).Methods("GET")
	r.HandleFunc("/{short}/", s.redirect).Methods("GET")
	return r
}

func (s *Server) Run(port string) error {
	r := s.router()
	return http.ListenAndServe(port, r)
}

func New(db storage.Storage, urlBase string, shortLength int) *Server {
	return &Server{db: db, urlBase: urlBase, shortLength: shortLength}
}
