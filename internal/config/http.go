package config

import (
	"time"

	"emperror.dev/errors"
)

type HTTP struct {
	HTTPHost                       string `envconfig:"GAMBLING_HTTP_HOST"`
	HTTPPort                       string `envconfig:"GAMBLING_HTTP_PORT"`
	HTTPReadHeaderTimeoutInSeconds int    `envconfig:"GAMBLING_HTTP_READ_HEADER_TIMEOUT_IN_SECONDS"`
	HTTPReadinessEndpoint          string `envconfig:"GAMBLING_HTTP_READINESS_ENDPOINT"`
}

func (h HTTP) validate() error {
	if h.HTTPPort == "" {
		return errors.New("empty GAMBLING_HTTP_PORT")
	}

	if h.HTTPReadHeaderTimeoutInSeconds <= 0 {
		return errors.New("GAMBLING_HTTP_READ_HEADER_TIMEOUT_IN_SECONDS is less or equal zero")
	}

	if h.HTTPReadinessEndpoint == "" {
		return errors.New("empty GAMBLING_HTTP_READINESS_ENDPOINT")
	}

	return nil
}

func (h HTTP) Host() string {
	return h.HTTPHost
}

func (h HTTP) Port() string {
	return h.HTTPPort
}

func (h HTTP) ReadHeaderTimeout() time.Duration {
	return time.Duration(h.HTTPReadHeaderTimeoutInSeconds) * time.Second
}

func (h HTTP) ReadinessEndpoint() string {
	return h.HTTPReadinessEndpoint
}
