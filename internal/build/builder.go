package build

import (
	"context"
	"net/http"

	"github.com/avstrong/gambling/internal/config"
	"github.com/avstrong/gambling/internal/gmanager"
	"github.com/avstrong/gambling/internal/uservice"
)

type Builder struct {
	config *config.Config

	uService *uservice.Service
	gManager *gmanager.Manager

	server *http.Server

	shutdown    shutdown
	healthcheck healthcheck
}

func New(_ context.Context, conf *config.Config) *Builder {
	b := Builder{config: conf} //nolint:exhaustruct

	return &b
}
