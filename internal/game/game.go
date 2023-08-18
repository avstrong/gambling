package game

import (
	"math/rand"
	"time"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/wallet"
	"github.com/google/uuid"
)

type CoinSide bool

const (
	CoinSideHeads CoinSide = true
	CoinSideTails CoinSide = false
)

type state string

const (
	stateWaitingForPlayers state = "stateWaitingForPlayers"
	stateWaitingToStart    state = "stateWaitingToStart"
	statePlaying           state = "statePlaying"
	stateFinished          state = "stateFinished"
)

type Player struct {
	UserID uuid.UUID
	Choice CoinSide
}

type Players []*Player

type Config struct {
	PlayerCount   int
	EntryFee      float64
	EntryCurrency wallet.Currency
}

type Game struct {
	id         uuid.UUID
	name       string
	config     *Config
	players    Players
	winners    Players
	resultSide CoinSide
	state      state
	totalBet   float64
	winAmount  float64
	createdAt  time.Time
	finishAt   time.Time
}

func New(id uuid.UUID, name string, config *Config) *Game {
	//nolint:exhaustruct
	return &Game{
		id:        id,
		name:      name,
		config:    config,
		state:     stateWaitingForPlayers,
		totalBet:  config.EntryFee * float64(config.PlayerCount),
		createdAt: time.Now().UTC(),
	}
}

func (g *Game) ID() uuid.UUID {
	return g.id
}

func (g *Game) Name() string {
	return g.name
}

func (g *Game) EntryFee() float64 {
	return g.config.EntryFee
}

func (g *Game) EntryCurrency() wallet.Currency {
	return g.config.EntryCurrency
}

func (g *Game) Winners() Players {
	return g.winners
}

func (g *Game) WinAmount() float64 {
	return g.winAmount
}

func (g *Game) FinishAt() time.Time {
	return g.finishAt
}

func (g *Game) AddPlayer(player *Player) error {
	if g.state != stateWaitingForPlayers {
		return errors.Errorf("cannot add player, game %v is in state %v", g.name, g.state)
	}

	g.players = append(g.players, player)

	if len(g.players) == g.config.PlayerCount {
		g.state = stateWaitingToStart
	}

	return nil
}

func (g *Game) Play() error {
	if g.state != stateWaitingToStart {
		return errors.Errorf("cannot play, game %v is in state %v", g.name, g.state)
	}

	g.state = statePlaying
	//nolint:gosec,gomnd
	g.resultSide = rand.Intn(2) == 0

	g.finish()

	return nil
}

func (g *Game) finish() {
	g.state = stateFinished

	for _, player := range g.players {
		if g.resultSide == player.Choice {
			g.winners = append(g.winners, player)
		}
	}

	g.winAmount = g.totalBet / float64(len(g.winners))
	g.finishAt = time.Now().UTC()
}
