package customhttp

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	JSON(w http.ResponseWriter, status int, payload interface{})
	Error(w http.ResponseWriter, status int, message string)
}

type responder struct{}

func NewResponder() Responder {
	return &responder{}
}

func (r *responder) JSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (r *responder) Error(w http.ResponseWriter, status int, message string) {
	r.JSON(w, status, map[string]string{"error": message})
}
