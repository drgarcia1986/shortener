package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/drgarcia1986/shortener/url"
)

type shortenerRequest struct {
	Url string `json:"url"`
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
	if err != nil || request.Url == "" {
		respondBadRequest(w)
		return
	}

	short := url.GenerateShort(s.shortLength)
	u := url.Url{Short: short, Original: request.Url}
	if err := s.db.Set(&u); err != nil {
		respondServerError(w)
		return
	}
	respondWith(w, http.StatusCreated, u)
}
