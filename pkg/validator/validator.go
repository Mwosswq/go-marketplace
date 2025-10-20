package validator

import (
	"errors"
	"regexp"

	"go.uber.org/zap"
)

type Validator interface {
	ValidatePassword(password string) error
	ValidateEmail(email string) error
}

type validator struct {
	logger *zap.Logger
}

func New(l *zap.Logger) Validator {
	return &validator{logger: l}
}

func (v *validator) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password too short")
	}
	regExp := regexp.MustCompile("^[A-Za-z0-9]+$")
	if !regExp.MatchString(password) {
		return errors.New("password must contain only letters and numbers")
	}
	return nil
}

func (v *validator) ValidateEmail(email string) error {
	regExp := regexp.MustCompile("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$")
	if !regExp.MatchString(email) {
		return errors.New("email must contain @")
	}

	return nil
}
