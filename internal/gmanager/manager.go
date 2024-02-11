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
	UpdateGame(ctx context.Context, g *game.Game) error
	GetGame(ctx context.Context, id uuid.UUID) (*game.Game, error)
	SavePlayer(ctx context.Context, player *Player) error
	UpdatePlayer(ctx context.Context, player *Player) error
	GetPlayer(ctx context.Context, id string) (*Player, error)
	IsAlreadyExistErr(err error) bool
}

type userManager interface {
	Withdraw(ctx context.Context, input *uservice.WithdrawInput) (*uservice.WithdrawOutput, error)
	GetUser(ctx context.Context, input *uservice.GetUserInput) (*uservice.GetUserOutput, error)
	IsNotFoundErr(err error) bool
}

type Player struct {
	userID      string
	wins        int
	lastWinTime time.Time
}

func (u *Player) UserID() string {
	return u.userID
}

func (u *Player) Wins() int {
	return u.wins
}

func (u *Player) LastWinTime() time.Time {
	return u.lastWinTime
}

type Manager struct {
	storage     storage
	userManager userManager
}

func New(storage storage, uManager userManager) *Manager {
	return &Manager{
		storage:     storage,
		userManager: uManager,
	}
}
