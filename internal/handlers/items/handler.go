package items

import (
	"context"
	"encoding/json"
	"main/internal/domain"
	"main/internal/services/items"
	responder "main/pkg"
	"net/http"

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

func (h *Handler) CreateItem(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var item domain.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.logger.Error("Error decode body: %v", zap.Error(err))
		return err
	}

	if err := h.service.Create(ctx, &item); err != nil {
		return err
	}

	if err := h.responder.Created(w, map[string]string{"message": "Success"}); err != nil {
		h.logger.Error("Error encode response %v", zap.Error(err))
		return err
	}

	return nil
}

func (h *Handler) GetAllItems(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	res, err := h.service.GetAllItems(ctx)
	if err != nil {
		return err
	}

	if err := h.responder.OK(w, map[string][]domain.Item{"data": res}); err != nil {
		h.logger.Error("Response error: ", zap.Error(err))
	}

	return nil
}
