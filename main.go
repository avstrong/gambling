package main

import (
	"context"
	"log"
	"os"
	"runtime/debug"

	"github.com/avstrong/gambling/internal/cmd"
	"github.com/avstrong/gambling/internal/config"
	"github.com/rs/zerolog"
)

//nolint:gochecknoglobals // build flag
var buildSHA = "local"

//nolint:gochecknoglobals // build flag
var buildTime = "local"

//nolint:gochecknoglobals // build flag
var buildUser = "local"

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalln("Cannot read build info")
	}

	zlog := zerolog.
		New(os.Stdout).
		Level(zerolog.Level(conf.App().LogLevel())).
		With().
		Timestamp().
		Stack().
		Caller().
		Str("app", conf.App().Name()).
		Str("buildTime", buildTime).
		Str("buildUser", buildUser).
		Str("buildSHA", buildSHA).
		Str("goVersion", buildInfo.GoVersion).
		Logger()

	ctx := zlog.WithContext(context.Background())

	var exitCode int

	if err = cmd.Run(ctx, conf, &zlog); err != nil {
		zlog.Err(err).Send()

		exitCode = 1
	}

	os.Exit(exitCode)
}
