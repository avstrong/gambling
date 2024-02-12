package config

import "emperror.dev/errors"

type GRPC struct {
	GRPCNetworkType string `envconfig:"GAMBLING_GRPC_NETWORK_TYPE"`
	GRPCHost        string `envconfig:"GAMBLING_GRPC_HOST"`
	GRPCPort        string `envconfig:"GAMBLING_GRPC_PORT"`
}

func (g GRPC) validate() error {
	if g.GRPCNetworkType == "" {
		return errors.New("empty GAMBLING_GRPC_NETWORK_TYPE")
	}

	if g.GRPCPort == "" {
		return errors.New("empty GAMBLING_GRPC_PORT")
	}

	return nil
}

func (g GRPC) NetworkType() string {
	return g.GRPCNetworkType
}

func (g GRPC) Host() string {
	return g.GRPCHost
}

func (g GRPC) Port() string {
	return g.GRPCPort
}
