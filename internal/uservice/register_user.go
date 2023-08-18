package uservice

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/user"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RegisterUserInput struct {
	Email string `validate:"required,email"`
}

func (i *RegisterUserInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate RegisterUserInput")
	}

	return nil
}

type RegisterUserOutput struct {
	UserID uuid.UUID
}

func (s *Service) RegisterUser(ctx context.Context, input *RegisterUserInput) (*RegisterUserOutput, error) {
	u := user.New(uuid.New(), input.Email)

	err := s.storage.SaveUser(ctx, u)
	if s.storage.IsAlreadyExistErr(err) {
		return nil, errors.Wrapf(ErrAlreadyExists, "save user %v", input.Email)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "save user %v", input.Email)
	}

	return &RegisterUserOutput{UserID: u.ID()}, nil
}
