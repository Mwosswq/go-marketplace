package items

import (
	"context"
	"items-service/internal/domain"
	"items-service/internal/services/items"

	pb "github.com/marketplace-go/contracts/items"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	pb.UnimplementedItemServiceServer
	service items.Service
}

func NewItemsHandler(service items.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	var item = &domain.Item{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
	}

	id, err := h.service.CreateItem(ctx, item)

	if err != nil {
		return nil, err
	}

	return &pb.CreateItemResponse{Id: id}, nil
}

func (h *Handler) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {

	item, err := h.service.GetItem(ctx, req.Id)

	if err != nil {
		return &pb.GetItemResponse{}, nil
	}

	return &pb.GetItemResponse{
		Id:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   timestamppb.New(item.CreatedAt),
		Price:       item.Price,
	}, nil
}
