package users

import (
	"context"
	"errors"
	"main/internal/domain"
	"main/internal/repository/users"
	"main/pkg/hash"
	"main/pkg/tokens"
	"main/pkg/validator"
	"sync"

	"go.uber.org/zap"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) error
	LoginUser(ctx context.Context, user domain.User) (map[string]string, error)
}

type service struct {
	repo      users.Repository
	logger    *zap.Logger
	validator validator.Validator
}

func NewUsersService(r users.Repository, l *zap.Logger, v validator.Validator) Service {
	return &service{repo: r, logger: l, validator: v}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) error {
	if err := s.validator.ValidatePassword(user.Password); err != nil {
		s.logger.Error("validation error: ", zap.Error(err))
		return err
	}
	if err := s.validator.ValidateEmail(user.Email); err != nil {
		s.logger.Error("validation error: ", zap.Error(err))
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	errs := make(chan error, 2)

	go func() {
		defer wg.Done()

		exists, err := s.repo.CheckEmailExisting(ctx, user.Email)

		if err != nil {
			errs <- err
			return
		}

		if exists {
			errs <- errors.New("email already exists")
			return
		}
	}()

	go func() {
		defer wg.Done()

		exists, err := s.repo.CheckUsernameExisting(ctx, user.Username)

		if err != nil {
			errs <- err
			return
		}

		if exists {
			errs <- errors.New("username already exists")
			return
		}
	}()

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			s.logger.Error("Can not create user: ", zap.Error(err))
			return err
		}
	}

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := s.repo.CreateUser(ctx, user); err != nil {
		s.logger.Error("Error while creating user: ", zap.Error(err))
		return err
	}

	return nil
}

func (s *service) LoginUser(ctx context.Context, user domain.User) (map[string]string, error) {
	if len(user.Username) == 0 {
		s.logger.Error("username is empty")
		return make(map[string]string), errors.New("missing fields")
	}
	if len(user.Password) == 0 {
		s.logger.Error("password is empty")
		return make(map[string]string), errors.New("missing fields")
	}

	u, err := s.repo.GetUserForLogin(ctx, user.Username)

	if err != nil {
		s.logger.Error("error while getting user: ", zap.Error(err))
		return make(map[string]string), err
	}

	if !hash.ComparePasswords(u.Password, user.Password) {
		s.logger.Error("invalid credentials")
		return make(map[string]string), errors.New("invalid credentials")
	}

	accessToken, err := tokens.SignAccessToken(u.ID)

	if err != nil {
		s.logger.Error("error while signing access token: ", zap.Error(err))
		return make(map[string]string), err
	}

	refreshToken, err := tokens.SignRefreshToken(u.ID)

	if err != nil {
		s.logger.Error("error while signing access token: ", zap.Error(err))
		return make(map[string]string), err
	}

	return map[string]string{"accessToken": accessToken, "refreshToken": refreshToken}, nil
}
