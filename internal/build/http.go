package build

import (
	"context"
	"net/http"

	"github.com/avstrong/gambling/internal/config"
	"github.com/avstrong/gambling/internal/transport/web"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func (b *Builder) HTTPServer(ctx context.Context, conf *config.Config, lg *zerolog.Logger) (*http.Server, error) {
	if b.server == nil {
		uService, err := b.UService(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "build user service")
		}

		gManager, err := b.GManager(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "build game manager")
		}

		srv, err := web.New(ctx, &web.Conf{
			ServerLogger:      nil,
			Host:              conf.HTTP().Host(),
			Port:              conf.HTTP().Port(),
			ReadHeaderTimeout: conf.HTTP().ReadHeaderTimeout(),
		}, uService, gManager, lg)
		if err != nil {
			return nil, errors.Wrap(err, "build http server")
		}

		router := srv.Router()
		router.HandleFunc(conf.HTTP().ReadinessEndpoint(), b.healthcheck.handler)

		server := srv.Srv()

		b.shutdown.addHiPriority(func(ctx context.Context) error {
			return errors.Wrap(server.Shutdown(ctx), "shutdown http server")
		})

		b.server = server
	}

	return b.server, nil
}
