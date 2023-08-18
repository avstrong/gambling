package user

import (
	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/google/uuid"
)

type User struct {
	id      uuid.UUID
	email   string
	wallets map[wallet.Currency]*wallet.Wallet
}

func New(id uuid.UUID, email string) *User {
	return &User{
		id:      id,
		email:   email,
		wallets: make(map[wallet.Currency]*wallet.Wallet),
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) getWallet(currency wallet.Currency) (*wallet.Wallet, error) {
	w, exists := u.wallets[currency]
	if exists {
		return w, nil
	}

	if currency != wallet.CurrencyUSD && currency != wallet.CurrencyEUR {
		return nil, errors.Errorf("unsupported currency: %s", currency)
	}

	// lazy init
	w = wallet.New(currency)
	u.wallets[currency] = w

	return w, nil
}

func (u *User) Deposit(amount float64, currency wallet.Currency) (float64, error) {
	w, err := u.getWallet(currency)
	if err != nil {
		return 0, errors.Wrapf(err, "get wallet %v", currency)
	}

	//nolint:wrapcheck
	return w.Add(amount)
}

func (u *User) Withdraw(amount float64, currency wallet.Currency) (float64, error) {
	w, err := u.getWallet(currency)
	if err != nil {
		return 0, errors.Wrapf(err, "get wallet %v", currency)
	}

	//nolint:wrapcheck
	return w.Subtract(amount)
}

func (u *User) Balance(currency wallet.Currency) (float64, error) {
	w, err := u.getWallet(currency)
	if err != nil {
		return 0, errors.Wrapf(err, "get wallet %v", currency)
	}

	return w.Balance(), nil
}

func (u *User) Transactions(currency wallet.Currency) (wallet.Transactions, error) {
	w, err := u.getWallet(currency)
	if err != nil {
		return nil, errors.Wrapf(err, "get wallet %v", currency)
	}

	return w.Transactions(), nil
}
