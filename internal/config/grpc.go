package config

type GRPC struct {
	GRPCNetworkType string `envconfig:"GRPC_NETWORK_TYPE"`
	GRPCPort        string `envconfig:"GRPC_PORT"`
}

func (h GRPC) NetworkType() string {
	return h.GRPCNetworkType
}

func (h GRPC) Port() string {
	return h.GRPCPort
}
