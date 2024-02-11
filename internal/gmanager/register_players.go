package gmanager

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/game"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type RegisterPlayerInput struct {
	GameID       uuid.UUID     `json:"gameID" validate:"required"` //nolint:tagliatelle
	UserID       string        `json:"userID" validate:"required"` //nolint:tagliatelle
	PlayerChoice game.CoinSide `json:"playerChoice"`
}

func (i *RegisterPlayerInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate RegisterPlayerInput")
	}

	return nil
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

	err = m.storage.SavePlayer(ctx, &Player{userID: input.UserID}) //nolint:exhaustruct // it's enough
	if m.storage.IsAlreadyExistErr(err) {
		return nil, errors.Wrap(ErrAlreadyExists, "save player")
	}

	if err != nil {
		return nil, errors.Wrapf(err, "save player to db with id %v", input.UserID)
	}

	_, err = m.userManager.Withdraw(ctx, &uservice.WithdrawInput{UserID: input.UserID, Amount: g.EntryFee(), Currency: g.EntryCurrency()})
	if err != nil {
		return nil, errors.Wrapf(
			errors.Wrap(ErrInsufficientFunds, err.Error()),
			"withdraw %v in %v from user %v to play game",
			g.EntryFee(),
			g.EntryCurrency(),
			input.UserID,
		)
	}

	return &RegisterPlayerOutput{}, nil
}
