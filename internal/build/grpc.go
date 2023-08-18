package build

import (
	"context"

	srv "github.com/avstrong/gambling/internal/transport/grpc"
	pb "github.com/avstrong/gambling/pkg/game"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func (b *Builder) GRPCServer(ctx context.Context, _ *zerolog.Logger) (*grpc.Server, error) {
	uService, err := b.UService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "build user service")
	}

	opts := []grpc.UnaryServerInterceptor{
		srv.RecoverUnaryServerInterceptor(ctx, true),
		srv.AccessLogUnaryServerInterceptor(ctx),
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(opts...))

	//nolint:exhaustruct
	pb.RegisterAPIServer(server, &srv.Server{
		UService: uService,
	})

	b.shutdown.addHiPriority(func(ctx context.Context) error {
		server.GracefulStop()

		return nil
	})

	return server, nil
}
