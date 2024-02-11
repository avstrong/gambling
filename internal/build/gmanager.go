package build

import (
	"context"

	"github.com/avstrong/gambling/internal/gmanager"
	"github.com/avstrong/gambling/internal/storage/memory"
)

func (b *Builder) GManager(_ context.Context) (*gmanager.Manager, error) {
	if b.gManager == nil {
		service := gmanager.New(memory.New(&memory.Config{ShardCount: 32}), b.uService) //nolint:gomnd

		b.gManager = service
	}

	return b.gManager, nil
}
