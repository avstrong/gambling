package converter

import (
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/avstrong/gambling/internal/wallet"
	pb "github.com/avstrong/gambling/pkg/game"
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
	return &uservice.BalanceInput{
		UserID:   input.GetUserId(),
		Currency: pbToCurrency(input.Currency),
	}, nil
}
