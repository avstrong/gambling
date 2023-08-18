package config

import "emperror.dev/errors"

type App struct {
	AppEnvironment string `envconfig:"APP_ENVIRONMENT"`
	AppName        string `envconfig:"APP_NAME"`
	AppLogLevel    int    `envconfig:"APP_LOG_LEVEL"`
}

func (a App) validate() error {
	allowedEnvironments := map[string]struct{}{
		"local": {},
		"stage": {},
		"prod":  {},
	}

	if a.AppEnvironment == "" {
		return errors.New("empty GAME_APP_ENVIRONMENT")
	}

	if _, ok := allowedEnvironments[a.AppEnvironment]; !ok {
		return errors.Errorf("allowed values for GAME_APP_ENVIRONMENT are %v", allowedEnvironments)
	}

	if a.AppName == "" {
		return errors.New("empty GAME_APP_NAME")
	}

	allowedLogLevels := map[int]string{
		-1: "TraceLevel",
		0:  "DebugLevel",
		1:  "InfoLevel",
		2:  "WarnLevel",
		3:  "ErrorLevel",
		4:  "FatalLevel",
		5:  "PanicLevel",
	}

	if _, ok := allowedLogLevels[a.AppLogLevel]; !ok {
		return errors.Errorf("allowed values for GAME_APP_LOG_LEVEL are %v", allowedEnvironments)
	}

	return nil
}

func (a App) Environment() string {
	return a.AppEnvironment
}

func (a App) Name() string {
	return a.AppName
}

func (a App) LogLevel() int {
	return a.AppLogLevel
}
