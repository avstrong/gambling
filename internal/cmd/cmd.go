package cmd

import (
	"context"
	"net"
	"net/http"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/build"
	"github.com/avstrong/gambling/internal/config"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, conf *config.Config, logger *zerolog.Logger) error {
	//nolint:exhaustruct // it's enough to use the struct
	root := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}

	root.AddCommand(
		start(ctx, conf, logger),
	)

	if err := root.ExecuteContext(ctx); err != nil {
		return errors.Wrap(err, "run application")
	}

	return nil
}

func start(ctx context.Context, conf *config.Config, _ *zerolog.Logger) *cobra.Command {
	//nolint:exhaustruct // it's enough to use the struct
	return &cobra.Command{
		Use:   "start",
		Short: "start app",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			builder := build.New(ctx, conf)

			grpcSrv, err := builder.GRPCServer(ctx, zerolog.Ctx(ctx))
			if err != nil {
				return errors.Wrap(err, "build grpc server")
			}

			reflection.Register(grpcSrv)

			srv, err := builder.HTTPServer(ctx, conf, zerolog.Ctx(ctx))
			if err != nil {
				return errors.Wrap(err, "build http server")
			}

			go func() {
				if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					zerolog.Ctx(ctx).Err(err).Msg("run http server")
				}
			}()

			go func() {
				builder.WaitShutdown(ctx)
				cancel()
			}()

			listener, err := net.Listen(conf.GRPC().NetworkType(), net.JoinHostPort(conf.GRPC().Host(), conf.GRPC().Port()))
			if err != nil {
				return errors.Wrap(err, "create network listener")
			}

			if err = grpcSrv.Serve(listener); !errors.Is(err, grpc.ErrServerStopped) {
				zerolog.Ctx(ctx).Err(errors.Wrap(err, "run grpc server")).Send()
			}

			<-ctx.Done()

			return nil
		},
	}
}
