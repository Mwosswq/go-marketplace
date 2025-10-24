package auth

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/handlers/users"
	"main/internal/repository/auth"
	"main/pkg/hash"
	"main/pkg/validator"

	"go.uber.org/zap"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) error
	LoginUser(ctx context.Context, user users.UserCredentials) (string, error)
}

type service struct {
	repo      auth.Repository
	logger    *zap.Logger
	validator validator.Validator
}

func NewAuthService(r auth.Repository, l *zap.Logger, v validator.Validator) Service {
	return &service{repo: r, logger: l, validator: v}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) error {
	if err := s.validator.ValidateEmail(user.Email); err != nil {
		s.logger.Error("email validation error: ", zap.Error(err))
		return errors.New("invalid email")
	}

	if err := s.validator.ValidatePassword(user.Password); err != nil {
		s.logger.Error("password validation error: ", zap.Error(err))
		return errors.New("invalid password")
	}

	hashedPassword, err := hash.HashPassword(user.Password)

	if err != nil {
		s.logger.Error("password hashing error: ", zap.Error(err))
		return errors.New("hashing error")
	}

	user.Password = hashedPassword

	if err := s.repo.CreateUser(ctx, user); err != nil {
		s.logger.Error("creating user error: ", zap.Error(err))
		return errors.New("creating error")
	}

	return nil
}

func (s *service) LoginUser(ctx context.Context, user users.UserCredentials) (string, error) {
	if user.Password == "" {
		s.logger.Error("password is empty")
		return "", errors.New("password is empty")
	}

	if user.Username == "" {
		s.logger.Error("username is empty")
		return "", errors.New("username is empty")
	}

	return "", nil
}
