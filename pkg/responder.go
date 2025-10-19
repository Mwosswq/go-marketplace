package responder

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Responder interface {
	OK(w http.ResponseWriter, data interface{})
	Created(w http.ResponseWriter, data interface{})
	Error(w http.ResponseWriter, data interface{})
}

type responder struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Responder {
	return &responder{logger: logger}
}

func (r *responder) OK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		r.logger.Error("Error while encoding response: ", zap.Error(err))
	}
}

func (r *responder) Created(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		r.logger.Error("Error while encoding response: ", zap.Error(err))
	}

}

func (r *responder) Error(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		r.logger.Error("Error while encoding response: ", zap.Error(err))
	}
}
