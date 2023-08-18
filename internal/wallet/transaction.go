package wallet

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType uint8

const (
	_                         TransactionType = 0
	TransactionTypeDeposit    TransactionType = 1
	TransactionTypeWithdrawal TransactionType = 2
)

type Transaction struct {
	id        uuid.UUID
	amount    float64
	action    TransactionType
	createdAt time.Time
}

func NewTransaction(id uuid.UUID, amount float64, action TransactionType, createdAt time.Time) *Transaction {
	return &Transaction{
		id:        id,
		amount:    amount,
		action:    action,
		createdAt: createdAt,
	}
}

type Transactions []*Transaction
