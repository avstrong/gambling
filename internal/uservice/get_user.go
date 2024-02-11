package uservice

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/user"
)

type GetUserInput struct {
	ID string
}

type GetUserOutput struct {
	User *user.User
}

func (s *Service) GetUser(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
	u, err := s.storage.GetUser(ctx, input.ID)
	if s.storage.IsNotFoundErr(err) {
		return nil, errors.Wrapf(ErrNotFound, "get user %v", input.ID)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "get user %v", input.ID)
	}

	return &GetUserOutput{User: u}, nil
}

func (s *Service) IsNotFoundErr(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrNotFound)
}
