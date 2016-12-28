package server

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

var serverError = errorResponse{Error: "Internal Server Error"}
var clientError = errorResponse{Error: "Bad Request"}
var NotFoundError = errorResponse{Error: "Not Found"}

func respondWith(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}

func respondServerError(w http.ResponseWriter) {
	respondWith(w, http.StatusInternalServerError, serverError)
}

func respondBadRequest(w http.ResponseWriter) {
	respondWith(w, http.StatusBadRequest, clientError)
}

func respondNotFound(w http.ResponseWriter) {
	respondWith(w, http.StatusNotFound, NotFoundError)
}
