package uservice

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BalanceInput struct {
	UserID   uuid.UUID       `validate:"required,uuid4"`
	Currency wallet.Currency `validate:"required"`
}

func (i *BalanceInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate BalanceInput")
	}

	kindSet := map[wallet.Currency]struct{}{
		wallet.CurrencyUSD: {},
		wallet.CurrencyEUR: {},
	}

	if _, ok := kindSet[i.Currency]; !ok {
		return errors.Wrapf(
			ErrInvalidFieldValue,
			"invalid Currency: '%v'. Currency must be one of the permitted currencies",
			i.Currency,
		)
	}

	return nil
}

type BalanceOutput struct {
	Balance float64
}

func (s *Service) Balance(ctx context.Context, input *BalanceInput) (*BalanceOutput, error) {
	u, err := s.storage.GetUser(ctx, input.UserID)
	if s.storage.IsNotFoundErr(err) {
		return nil, errors.Wrapf(ErrNotFound, "get user %v", input.UserID)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "get user %v", input.UserID)
	}

	balance, err := u.Balance(input.Currency)
	if err != nil {
		return nil, errors.Wrapf(err, "get balance of wallet %v", input.Currency)
	}

	return &BalanceOutput{Balance: balance}, nil
}
