package gmanager

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/game"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateGameInput struct {
	Name          string          `json:"name" validate:"required"`
	MaxPlayers    int             `json:"maxPlayers" validate:"required,gte=1"`
	EntryFee      float64         `json:"entryFee" validate:"required,gte=1"`
	EntryCurrency wallet.Currency `json:"entryCurrency" validate:"required"`
}

func (i *CreateGameInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate CreateGameInput")
	}

	kindSet := map[wallet.Currency]struct{}{
		wallet.CurrencyUSD: {},
		wallet.CurrencyEUR: {},
	}

	if _, ok := kindSet[i.EntryCurrency]; !ok {
		return errors.Wrapf(
			ErrInvalidFieldValue,
			"invalid currency: '%v'. currency must be one of the permitted currencies",
			i.EntryCurrency,
		)
	}

	return nil
}

type CreateGameOutput struct {
	ID uuid.UUID `json:"id"`
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
