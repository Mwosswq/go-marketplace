package items

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/repository/items"
	"main/pkg/app_error"

	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, item *domain.Item) error
	GetAllItems(ctx context.Context) ([]domain.Item, error)
	RemoveItem(ctx context.Context, id int) error
}

type service struct {
	repo   items.Repository
	logger *zap.Logger
}

func NewItemService(repo items.Repository, logger *zap.Logger) Service {
	return &service{repo: repo, logger: logger}
}

func (s *service) Create(ctx context.Context, item *domain.Item) error {
	if err := s.validateItem(item); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, item); err != nil {
		return err
	}

	return nil
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

func (s *service) GetAllItems(ctx context.Context) ([]domain.Item, error) {
	res, err := s.repo.GetAllItems(ctx)
	if err != nil {
		s.logger.Error("Error while getting items list:", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *service) RemoveItem(ctx context.Context, id int) error {
	if id == 0 {
		s.logger.Error("Id can not be zero value")
		return errors.New("ID is zero value")
	}

	if err := s.repo.RemoveItem(ctx, id); err != nil {
		s.logger.Error("Error while removing item: ", zap.Error(err))
		return err
	}

	return nil
}
