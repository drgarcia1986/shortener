package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/drgarcia1986/shortener/url"
)

type shortenerRequest struct {
	URL string `json:"url"`
}

func (s *Server) shortener(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondBadRequest(w)
		return
	}

	var request shortenerRequest
	err = json.Unmarshal(body, &request)
	if err != nil || request.URL == "" {
		respondBadRequest(w)
		return
	}

	short := url.GenerateShort(s.shortLength)
	u := url.URL{Short: short, Original: request.URL}
	if err := s.db.Set(&u); err != nil {
		respondServerError(w)
		return
	}
	respondWith(w, http.StatusCreated, u)
}
