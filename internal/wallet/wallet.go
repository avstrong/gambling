package wallet

import (
	"sync"
	"time"

	"emperror.dev/errors"
	"github.com/google/uuid"
)

type Currency string

const (
	CurrencyUnknown Currency = "Unknown"
	CurrencyUSD     Currency = "USD"
	CurrencyEUR     Currency = "EUR"
)

type Wallet struct {
	mu           sync.Mutex
	balance      float64
	currency     Currency
	transactions Transactions
}

func New(currency Currency) *Wallet {
	//nolint:exhaustruct
	return &Wallet{
		balance:  0,
		currency: currency,
		mu:       sync.Mutex{},
	}
}

func (w *Wallet) addTransaction(trxType TransactionType, amount float64) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.transactions = append(w.transactions, NewTransaction(uuid.New(), amount, trxType, time.Now().UTC()))
}

func (w *Wallet) Balance() float64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.balance
}

func (w *Wallet) Add(amount float64) (float64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if amount < 0 {
		return 0, errors.Errorf("amount must be positive: you provided %v", amount)
	}

	w.balance += amount

	w.addTransaction(TransactionTypeDeposit, amount)

	return w.balance, nil
}

func (w *Wallet) Subtract(amount float64) (float64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if amount < 0 {
		return 0, errors.Errorf("amount must be positive: you provided %v", amount)
	}

	if amount > w.balance {
		return 0, errors.New("insufficient funds")
	}

	w.balance -= amount

	w.addTransaction(TransactionTypeWithdrawal, amount)

	return w.balance, nil
}

func (w *Wallet) Transactions() Transactions {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.transactions
}
