//nolint:dupl
package uservice

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type WithdrawInput struct {
	UserID   uuid.UUID       `validate:"required,uuid4"`
	Amount   float64         `validate:"required,gte=0"`
	Currency wallet.Currency `validate:"required"`
}

func (i *WithdrawInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate TransactionsInput")
	}

	kindSet := map[wallet.Currency]struct{}{
		wallet.CurrencyUSD: {},
		wallet.CurrencyEUR: {},
	}

	if _, ok := kindSet[i.Currency]; !ok {
		return errors.Wrapf(ErrInvalidFieldValue, "invalid Currency: '%v'. Currency must be one of the permitted currencies", i.Currency)
	}

	return nil
}

type WithdrawOutput struct {
	Balance float64
}

func (s *Service) Withdraw(ctx context.Context, input *WithdrawInput) (*WithdrawOutput, error) {
	u, err := s.storage.GetUser(ctx, input.UserID)
	if s.storage.IsNotFoundErr(err) {
		return nil, errors.Wrapf(ErrNotFound, "get user %v", input.UserID)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "get user %v", input.UserID)
	}

	balance, err := u.Withdraw(input.Amount, input.Currency)
	if err != nil {
		return nil, errors.Wrapf(err, "take out %v amount from the account", input.Amount)
	}

	if err = s.storage.SaveUser(ctx, u); err != nil {
		return nil, errors.Wrapf(err, "save user balance %v", u.ID())
	}

	return &WithdrawOutput{Balance: balance}, nil
}
