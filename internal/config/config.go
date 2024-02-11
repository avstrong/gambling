package config

import (
	"log"

	"emperror.dev/errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	app  App
	http HTTP
	grpc GRPC
}

func New(filenames ...string) (*Config, error) {
	if len(filenames) != 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return nil, errors.Wrapf(err, "read %v files", filenames)
		}
	} else {
		if err := godotenv.Load(); err != nil {
			log.Println("Cannot load env file:", err.Error())
		}
	}

	var conf Config

	const appPrefix = ""

	if err := envconfig.Process(appPrefix, &conf.app); err != nil {
		return nil, errors.Wrap(err, "parse app envs")
	}

	if err := envconfig.Process(appPrefix, &conf.http); err != nil {
		return nil, errors.Wrap(err, "parse http envs")
	}

	if err := envconfig.Process(appPrefix, &conf.grpc); err != nil {
		return nil, errors.Wrap(err, "parse grpc envs")
	}

	if err := conf.validate(); err != nil {
		return nil, errors.Wrap(err, "validate config")
	}

	return &conf, nil
}

func (c Config) validate() error {
	if err := c.app.validate(); err != nil {
		return errors.Wrap(err, "validate app config")
	}

	if err := c.http.validate(); err != nil {
		return errors.Wrap(err, "validate http config")
	}

	if err := c.grpc.validate(); err != nil {
		return errors.Wrap(err, "validate grpc config")
	}

	return nil
}

func (c Config) App() App {
	return c.app
}

func (c Config) HTTP() HTTP {
	return c.http
}

func (c Config) GRPC() GRPC {
	return c.grpc
}
