package grpc

import (
	"context"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/transport/grpc/converter"
	"github.com/avstrong/gambling/internal/uservice"
	pb "github.com/avstrong/gambling/pkg/game"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAPIServer
	UService *uservice.Service
}

func New(uService *uservice.Service) *Server {
	//nolint:exhaustruct
	return &Server{
		UService: uService,
	}
}

func (s *Server) RetrieveBalance(ctx context.Context, input *pb.RetrieveBalanceInput) (*pb.RetrieveBalanceOutput, error) {
	balanceInput, err := converter.PBToBalanceInput(input)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, errors.Wrap(err, "convert pb to balanceInput").Error())
	}

	if err = balanceInput.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, errors.Wrap(err, "validate balanceInput").Error())
	}

	output, err := s.UService.Balance(ctx, balanceInput)
	if errors.Is(err, uservice.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "user with id %v not found", balanceInput.UserID.String())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, errors.Wrap(err, "retrieve balance").Error())
	}

	return converter.BalanceInputToPB(output), nil
}
