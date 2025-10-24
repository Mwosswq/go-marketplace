package users

import (
	"main/internal/services/users"
	"main/pkg/responder"
)

type Handler struct {
	service   users.Service
	responder responder.Responder
}

func NewUserHandler(s users.Service, r responder.Responder) *Handler {
	return &Handler{service: s, responder: r}
}
