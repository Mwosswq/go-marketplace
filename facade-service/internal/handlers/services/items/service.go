package items

import (
	"context"
	"facade-service/pkg/clients/items"

	pb "github.com/marketplace-go/contracts/items"
	"go.uber.org/zap"
)

type Service interface {
	CreateItem(ctx context.Context, item *pb.CreateItemRequest) (int32, error)
	GetItem(ctx context.Context, id int32) (*pb.GetItemResponse, error)
}

type service struct {
	client *items.Client
	logger *zap.Logger
}

func NewItemsService(c *items.Client, l *zap.Logger) Service {
	return &service{client: c, logger: l}
}

func (s *service) CreateItem(ctx context.Context, item *pb.CreateItemRequest) (int32, error) {
	s.logger.Info("requesting items service")
	resp, err := s.client.CreateItem(ctx, item)

	if err != nil {
		s.logger.Error("Error while creating user: ", zap.Error(err))
		return 0, err
	}

	return resp.GetId(), nil
}

func (s *service) GetItem(ctx context.Context, id int32) (*pb.GetItemResponse, error) {
	s.logger.Info("requesting items service")
	resp, err := s.client.GetItem(ctx, id)

	if err != nil {
		s.logger.Error("error while  getting item: ", zap.Error(err))
		return &pb.GetItemResponse{}, nil
	}

	return resp, nil
}
