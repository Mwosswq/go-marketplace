package responder

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	OK(w http.ResponseWriter, data any)
	Created(w http.ResponseWriter, id int32)
	Error(w http.ResponseWriter, code int, msg string)
}

type responder struct{}

func New() Responder {
	return &responder{}
}

func (r *responder) OK(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := map[string]any{"status": "ok", "data": data}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "error encoding response", 500)
		return
	}
}

func (r *responder) Created(w http.ResponseWriter, id int32) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := map[string]any{"status": "created", "id": id}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "error encoding response", 500)
		return
	}
}

func (r *responder) Error(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := map[string]any{"status": "error", "message": msg}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "error encoding response", 500)
	}
}
