package converter

import (
	"github.com/avstrong/gambling/internal/uservice"
	pb "github.com/avstrong/gambling/pkg/game"
)

func BalanceInputToPB(input *uservice.BalanceOutput) *pb.RetrieveBalanceOutput {
	return &pb.RetrieveBalanceOutput{Balance: input.Balance}
}
