package gmanager

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/game"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/google/uuid"
)

type RegisterPlayerInput struct {
	GameID       uuid.UUID
	UserID       uuid.UUID
	PlayerChoice game.CoinSide
}

type RegisterPlayerOutput struct{}

func (m *Manager) RegisterPlayer(ctx context.Context, input *RegisterPlayerInput) (*RegisterPlayerOutput, error) {
	g, err := m.storage.GetGame(ctx, input.GameID)
	if err != nil {
		return nil, errors.Wrapf(err, "get game %v", input.GameID)
	}

	if err = g.AddPlayer(&game.Player{UserID: input.UserID, Choice: input.PlayerChoice}); err != nil {
		return nil, errors.Wrapf(err, "add player %v to game %v with name", input.GameID, g.Name())
	}

	_, err = m.userManager.Withdraw(ctx, &uservice.WithdrawInput{UserID: input.UserID, Amount: g.EntryFee(), Currency: g.EntryCurrency()})
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"withdraw %v in %v from user %v to play game",
			g.EntryFee(),
			g.EntryCurrency(),
			input.UserID,
		)
	}

	return &RegisterPlayerOutput{}, nil
}
