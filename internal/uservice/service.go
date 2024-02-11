package uservice

import (
	"context"

	"github.com/avstrong/gambling/internal/user"
)

type storage interface {
	SaveUser(ctx context.Context, u *user.User) error
	UpdateUser(ctx context.Context, u *user.User) error
	GetUser(ctx context.Context, email string) (*user.User, error)
	IsAlreadyExistErr(err error) bool
	IsNotFoundErr(err error) bool
}

type Service struct {
	storage storage
}

func New(repo storage) *Service {
	return &Service{storage: repo}
}
