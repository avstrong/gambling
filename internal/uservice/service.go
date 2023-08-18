package uservice

import (
	"context"

	"github.com/avstrong/gambling/internal/user"
	"github.com/google/uuid"
)

type storage interface {
	SaveUser(ctx context.Context, u *user.User) error
	GetUser(ctx context.Context, id uuid.UUID) (*user.User, error)
	IsAlreadyExistErr(err error) bool
	IsNotFoundErr(err error) bool
}

type Service struct {
	storage storage
}

func New(repo storage) *Service {
	return &Service{storage: repo}
}
