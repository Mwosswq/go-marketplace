package items

import (
	"encoding/json"
	"facade-service/internal/handlers/services/items"
	"facade-service/pkg/responder"
	"net/http"

	pb "github.com/marketplace-go/contracts/items"
)

type Handler struct {
	service   items.Service
	responder responder.Responder
}

func NewItemsHandler(s items.Service, r responder.Responder) *Handler {
	return &Handler{service: s, responder: r}
}

func (h *Handler) CreatingItemRequest(w http.ResponseWriter, r *http.Request) {
	var item CreatingItemsRequest

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.responder.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	id, err := h.service.CreateItem(r.Context(), &pb.CreateItemRequest{
		Title:       item.Title,
		Description: item.Description,
		Price:       item.Price,
	})

	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, "server error")
	}

	h.responder.Created(w, id)
}
