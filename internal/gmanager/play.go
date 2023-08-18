package gmanager

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/google/uuid"
)

type StartGameInput struct {
	ID uuid.UUID
}

type StartGameOutput struct{}

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
			return nil, errors.Wrapf(err, "get user %v", winner.UserID)
		}

		p.Wins++
		p.LastWinTime = g.FinishAt()

		if err := m.storage.SavePlayer(ctx, p); err != nil {
			return nil, errors.Wrapf(err, "save player %v", winner.UserID)
		}
	}

	if err = m.storage.SaveGame(ctx, g); err != nil {
		return nil, errors.Wrapf(err, "save game %v with name %v", input.ID, g.Name())
	}

	return &StartGameOutput{}, nil
}
