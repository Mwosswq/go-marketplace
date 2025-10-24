package items

import (
	"encoding/json"
	"main/internal/domain"
	"main/internal/services/items"
	"main/pkg/responder"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type Handler struct {
	service   items.Service
	logger    *zap.Logger
	responder responder.Responder
}

func NewItemsHandler(service items.Service, logger *zap.Logger, responder responder.Responder) *Handler {
	return &Handler{service: service, logger: logger, responder: responder}
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item domain.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.logger.Error("Error decode body: %v", zap.Error(err))
		return
	}

	if err := h.service.Create(r.Context(), &item); err != nil {
		return
	}

	h.responder.Created(w, nil)
}

func (h *Handler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAllItems(r.Context())
	if err != nil {
		h.responder.Error(w, err)
		return
	}

	h.responder.OK(w, map[string][]domain.Item{"data": res})
}

func (h *Handler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		h.responder.Error(w, err)
		return
	}

	if err := h.service.RemoveItem(r.Context(), id); err != nil {
		h.responder.Error(w, err)
		return
	}

	h.responder.OK(w, map[string]string{"message": "successfuly removed"})
}
