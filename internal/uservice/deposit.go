//nolint:dupl
package uservice

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type DepositInput struct {
	UserID   uuid.UUID       `validate:"required,uuid4"`
	Amount   float64         `validate:"required,gte=0"`
	Currency wallet.Currency `validate:"required"`
}

func (i *DepositInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate DepositInput")
	}

	kindSet := map[wallet.Currency]struct{}{
		wallet.CurrencyUSD: {},
		wallet.CurrencyEUR: {},
	}

	if _, ok := kindSet[i.Currency]; !ok {
		return errors.Wrapf(
			ErrInvalidFieldValue,
			"invalid currency: '%v'. currency must be one of the permitted currencies",
			i.Currency,
		)
	}

	return nil
}

type DepositOutput struct {
	Balance float64
}

func (s *Service) Deposit(ctx context.Context, input *DepositInput) (*DepositOutput, error) {
	u, err := s.storage.GetUser(ctx, input.UserID)
	if s.storage.IsNotFoundErr(err) {
		return nil, errors.Wrapf(ErrNotFound, "get user %v", input.UserID)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "get user %v", input.UserID)
	}

	balance, err := u.Deposit(input.Amount, input.Currency)
	if err != nil {
		return nil, errors.Wrapf(err, "deposit %v amount into an account", input.Amount)
	}

	if err = s.storage.SaveUser(ctx, u); err != nil {
		return nil, errors.Wrapf(err, "save user balance %v", u.ID())
	}

	return &DepositOutput{Balance: balance}, nil
}
