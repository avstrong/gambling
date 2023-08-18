package build

import (
	"context"

	"github.com/avstrong/gambling/internal/storage/memory"
	"github.com/avstrong/gambling/internal/uservice"
)

func (b *Builder) UService(_ context.Context) (*uservice.Service, error) {
	if b.uService == nil {
		service := uservice.New(memory.New(&memory.Config{ShardCount: 32})) //nolint:gomnd

		b.uService = service
	}

	return b.uService, nil
}
