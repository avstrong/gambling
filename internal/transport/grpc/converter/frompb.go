package converter

import (
	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/avstrong/gambling/internal/wallet"
	pb "github.com/avstrong/gambling/pkg/game"
	"github.com/google/uuid"
)

func pbToCurrency(input pb.RetrieveBalanceInput_Currency) wallet.Currency {
	switch input {
	case pb.RetrieveBalanceInput_CURRENCY_USD:
		return wallet.CurrencyUSD
	case pb.RetrieveBalanceInput_CURRENCY_EUR:
		return wallet.CurrencyEUR
	case pb.RetrieveBalanceInput_CURRENCY_UNKNOWN:
		return wallet.CurrencyUnknown
	default:
		return wallet.CurrencyUnknown
	}
}

func PBToBalanceInput(input *pb.RetrieveBalanceInput) (*uservice.BalanceInput, error) {
	userID, err := uuid.FromBytes(input.GetUserId())
	if err != nil {
		return nil, errors.Wrap(err, "parse user id")
	}

	return &uservice.BalanceInput{
		UserID:   userID,
		Currency: pbToCurrency(input.Currency),
	}, nil
}
