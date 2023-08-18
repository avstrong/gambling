package config_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/avstrong/gambling/internal/config"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"
)

type Suite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite))
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func (s *Suite) TestNew_MustReadAllDataNoErr() {
	envs := map[string]string{
		"GAME_APP_ENVIRONMENT": "local",
		"GAME_APP_NAME":        "app",
		"GAME_APP_LOG_LEVEL":   "2",

		"GAME_HTTP_HOST": "localhost",
		"GAME_HTTP_PORT": "3434",
		"GAME_HTTP_READ_HEADER_TIMEOUT_IN_SECONDS": "10",
	}

	for name, value := range envs {
		s.Require().NoError(os.Setenv(name, value))
	}

	conf, err := config.New()
	s.Require().NoError(err)

	s.Run("app_envs", func() {
		s.Require().Equal(envs["GAME_APP_ENVIRONMENT"], conf.App().Environment())
		s.Require().Equal(envs["GAME_APP_NAME"], conf.App().Name())
		intValue, err := strconv.ParseInt(envs["GAME_APP_LOG_LEVEL"], 10, 8)
		s.Require().NoError(err)
		s.Require().Equal(int(intValue), conf.App().LogLevel())
	})

	s.Run("http_envs", func() {
		s.Require().Equal(envs["GAME_HTTP_HOST"], conf.HTTP().Host())
		s.Require().Equal(envs["GAME_HTTP_PORT"], conf.HTTP().Port())

		intValue, err := strconv.ParseInt(envs["GAME_HTTP_READ_HEADER_TIMEOUT_IN_SECONDS"], 10, 8)
		s.Require().NoError(err)
		s.Require().Equal(time.Duration(int(intValue))*time.Second, conf.HTTP().ReadHeaderTimeout())
	})

	s.T().Cleanup(func() {
		for name := range envs {
			s.Require().NoError(os.Unsetenv(name))
		}
	})
}
