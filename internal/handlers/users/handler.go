package users

import (
	"encoding/json"
	"main/internal/domain"
	"main/internal/services/users"
	"main/pkg/responder"
	"net/http"
)

type Handler struct {
	service   users.Service
	responder responder.Responder
}

func NewUserHandler(s users.Service, r responder.Responder) *Handler {
	return &Handler{service: s, responder: r}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.responder.Error(w, err)
		return
	}

	if err := h.service.CreateUser(r.Context(), user); err != nil {
		h.responder.Error(w, err)
		return
	}

	h.responder.Created(w, nil)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.responder.Error(w, err)
		return
	}

	u, err := h.service.LoginUser(r.Context(), user)

	if err != nil {
		h.responder.Error(w, err)
		return
	}

	h.responder.OK(w, u)
}
