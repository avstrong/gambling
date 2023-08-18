package gmanager

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/game"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/google/uuid"
)

type CreateGameInput struct {
	Name          string
	MaxPlayers    int
	EntryFee      float64
	EntryCurrency wallet.Currency
}

type CreateGameOutput struct {
	ID uuid.UUID
}

func (m *Manager) CreateGame(ctx context.Context, input *CreateGameInput) (*CreateGameOutput, error) {
	id := uuid.New()

	g := game.New(id, input.Name, &game.Config{
		PlayerCount:   input.MaxPlayers,
		EntryFee:      input.EntryFee,
		EntryCurrency: input.EntryCurrency,
	})

	if err := m.storage.SaveGame(ctx, g); err != nil {
		return nil, errors.Wrapf(err, "save game %v with name %v", id, input.Name)
	}

	return &CreateGameOutput{ID: id}, nil
}
