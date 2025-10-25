package items

import (
	"context"
	"items-service/internal/domain"
	"items-service/internal/repository/items"
	"items-service/pkg/app_error"

	"go.uber.org/zap"
)

type Service interface {
	CreateItem(ctx context.Context, item *domain.Item) (int32, error)
	GetItem(ctx context.Context, id int32) (domain.Item, error)
}

type service struct {
	repo   items.Repository
	logger *zap.Logger
}

func NewItemService(repo items.Repository, logger *zap.Logger) Service {
	return &service{repo: repo, logger: logger}
}

func (s *service) CreateItem(ctx context.Context, item *domain.Item) (int32, error) {
	if err := s.validateItem(item); err != nil {
		return 0, err
	}

	id, err := s.repo.CreateItem(ctx, item)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) validateItem(item *domain.Item) error {
	if item.Title == "" {
		s.logger.Error("Title can nt be empty")
		return app_error.ErrValidation
	}

	if item.Price <= 0 {
		s.logger.Error("Price can nt be zero")
		return app_error.ErrValidation
	}

	return nil
}

func (s *service) GetItem(ctx context.Context, id int32) (domain.Item, error) {
	res, err := s.repo.GetItem(ctx, id)
	if err != nil {
		s.logger.Error("Error while getting items list:", zap.Error(err))
		return domain.Item{}, err
	}

	return res, nil
}
