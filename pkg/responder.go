package responder

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	OK(w http.ResponseWriter, data interface{}) error
	Created(w http.ResponseWriter, data interface{}) error
}

type responder struct {
}

func New() Responder {
	return &responder{}
}

func (r *responder) OK(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	return nil
}

func (r *responder) Created(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	return nil
}
