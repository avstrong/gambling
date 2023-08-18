package build

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
)

func (b *Builder) ServerWaitShutdown(ctx context.Context, errChan chan error) {
	stopChan := b.getStopSignal()

	select {
	case s := <-stopChan:
		zerolog.Ctx(ctx).Info().Msgf("got %s os signal. application will be stopped", s)
	case <-errChan:
	}

	b.shutdown.do(ctx)
}

func (b *Builder) WaitShutdown(ctx context.Context) {
	stopSignals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	s := make(chan os.Signal, len(stopSignals))
	signal.Notify(s, stopSignals...)
	zerolog.Ctx(ctx).Info().Msgf("got %s os signal. application will be stopped", <-s)

	b.shutdown.do(ctx)
}

func (b *Builder) getStopSignal() chan os.Signal {
	stopSignals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	s := make(chan os.Signal, len(stopSignals))
	signal.Notify(s, stopSignals...)

	return s
}

type shutdownFn func(context.Context) error

type shutdown [3][]shutdownFn

func (s *shutdown) addHiPriority(fn shutdownFn) {
	s[0] = append(s[0], fn)
}

//nolint:unused
func (s *shutdown) addNormalPriority(fn shutdownFn) {
	s[1] = append(s[1], fn)
}

//nolint:unused
func (s *shutdown) addLowPriority(fn shutdownFn) {
	s[2] = append(s[2], fn)
}

func (s *shutdown) do(ctx context.Context) {
	for _, priorityShutdown := range s {
		for _, fn := range priorityShutdown {
			if err := fn(ctx); err != nil {
				zerolog.Ctx(ctx).Err(err).Send()
			}
		}
	}
}
