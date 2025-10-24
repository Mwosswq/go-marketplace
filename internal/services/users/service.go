package users

import (
	"main/internal/repository/users"
	"main/pkg/validator"

	"go.uber.org/zap"
)

type Service interface {
}

type service struct {
	repo      users.Repository
	logger    *zap.Logger
	validator validator.Validator
}

func NewUsersService(r users.Repository, l *zap.Logger, v validator.Validator) Service {
	return &service{repo: r, logger: l, validator: v}
}
