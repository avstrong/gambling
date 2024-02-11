package uservice

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/go-playground/validator/v10"
)

type TransactionsInput struct {
	UserID   string          `validate:"required,email"`
	Currency wallet.Currency `validate:"required"`
}

func (i *TransactionsInput) Validate() error {
	validate := validator.New()

	if err := validate.Struct(i); err != nil {
		return errors.Wrap(err, "validate TransactionsInput")
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

type TransactionsOutput struct {
	Transactions wallet.Transactions
}

func (s *Service) Transactions(ctx context.Context, input *TransactionsInput) (*TransactionsOutput, error) {
	u, err := s.storage.GetUser(ctx, input.UserID)
	if s.storage.IsNotFoundErr(err) {
		return nil, errors.Wrapf(ErrNotFound, "get user %v", input.UserID)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "get user %v", input.UserID)
	}

	transactions, err := u.Transactions(input.Currency)
	if err != nil {
		return nil, errors.Wrapf(err, "get transactions of wallet %v", input.Currency)
	}

	return &TransactionsOutput{Transactions: transactions}, nil
}
