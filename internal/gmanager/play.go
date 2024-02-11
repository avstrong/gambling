package gmanager

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/game"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type StartGameInput struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func (i *StartGameInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate StartGameInput")
	}

	return nil
}

type StartGameOutput struct {
	Winners game.Players `json:"winners"`
}

func (m *Manager) Play(ctx context.Context, input *StartGameInput) (*StartGameOutput, error) {
	g, err := m.storage.GetGame(ctx, input.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "get game %v", input.ID)
	}

	if err = g.Play(); err != nil {
		return nil, errors.Wrapf(err, "play game %v with name %v", input.ID, g.Name())
	}

	winAmount := g.WinAmount()
	winners := g.Winners()

	for _, winner := range winners {
		u, err := m.userManager.GetUser(ctx, &uservice.GetUserInput{ID: winner.UserID})
		if err != nil {
			// TODO
			return nil, errors.Wrapf(err, "get user %v", winner.UserID)
		}

		if _, err = u.User.Deposit(winAmount, g.EntryCurrency()); err != nil {
			return nil, errors.Wrapf(err, "deposit %v to user %v", winAmount, winner.UserID)
		}

		p, err := m.storage.GetPlayer(ctx, winner.UserID)
		if m.userManager.IsNotFoundErr(err) {
			// TODO
			continue
		}

		if err != nil {
			// TODO
			return nil, errors.Wrapf(err, "get player %v", winner.UserID)
		}

		p.wins++
		p.lastWinTime = g.FinishAt()

		if err := m.storage.UpdatePlayer(ctx, p); err != nil {
			return nil, errors.Wrapf(err, "update player %v", winner.UserID)
		}
	}

	if err = m.storage.UpdateGame(ctx, g); err != nil {
		return nil, errors.Wrapf(err, "update game %v with name %v", input.ID, g.Name())
	}

	return &StartGameOutput{
		Winners: winners,
	}, nil
}
