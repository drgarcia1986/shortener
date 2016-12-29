package server

import (
	"net/http"

	"github.com/drgarcia1986/shortener/url"
	"github.com/gorilla/mux"
)

func (s *Server) stats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	u, err := s.db.Get(short)
	if err != nil {
		if err == url.ErrNotFound {
			respondNotFound(w)
		} else {
			respondServerError(w)
		}
		return
	}
	respondWith(w, http.StatusOK, u)
}
