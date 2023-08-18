package gmanager

import (
	"context"
	"time"

	"github.com/avstrong/gambling/internal/game"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/google/uuid"
)

type storage interface {
	SaveGame(ctx context.Context, g *game.Game) error
	GetGame(ctx context.Context, id uuid.UUID) (*game.Game, error)
	SavePlayer(ctx context.Context, player *Player) error
	GetPlayer(ctx context.Context, id uuid.UUID) (*Player, error)
}

type userManager interface {
	Withdraw(ctx context.Context, input *uservice.WithdrawInput) (*uservice.WithdrawOutput, error)
	GetUser(ctx context.Context, input *uservice.GetUserInput) (*uservice.GetUserOutput, error)
	IsNotFoundErr(err error) bool
}

type Player struct {
	UserID      string
	Wins        int
	LastWinTime time.Time
}

type Manager struct {
	storage     storage
	userManager userManager
}
