package memory

import (
	"context"
	"hash/fnv"
	"sync"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/game"
	"github.com/avstrong/gambling/internal/gmanager"
	"github.com/avstrong/gambling/internal/user"
	"github.com/google/uuid"
)

var ErrUserExists = errors.New("user already exists")

var ErrUserNotExists = errors.New("user not found")

var ErrGameExists = errors.New("game already exists")

var ErrGameNotExists = errors.New("game not found")

var ErrPlayerExists = errors.New("player already exists")

var ErrPlayerNotExists = errors.New("player not found")

type memoryShard struct {
	mu         sync.Mutex
	UserData   map[string]*user.User
	GameData   map[uuid.UUID]*game.Game
	PlayerData map[string]*gmanager.Player
}

type InMemoryStorage struct {
	shards []*memoryShard
}

type Config struct {
	ShardCount int32
}

func New(config *Config) *InMemoryStorage {
	shards := make([]*memoryShard, config.ShardCount)
	for i := 0; i < int(config.ShardCount); i++ {
		shards[i] = &memoryShard{
			UserData:   make(map[string]*user.User),
			GameData:   make(map[uuid.UUID]*game.Game),
			PlayerData: make(map[string]*gmanager.Player),
			mu:         sync.Mutex{},
		}
	}

	return &InMemoryStorage{
		shards: shards,
	}
}

func (s *InMemoryStorage) getShard(id string) (*memoryShard, error) {
	hasher := fnv.New32()

	_, err := hasher.Write([]byte(id))
	if err != nil {
		return nil, errors.Wrap(err, "write hash")
	}

	return s.shards[hasher.Sum32()%uint32(len(s.shards))], nil
}

func (s *InMemoryStorage) SaveUser(_ context.Context, u *user.User) error {
	shard, err := s.getShard(u.Email())
	if err != nil {
		return errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.UserData[u.Email()]; exists {
		return ErrUserExists
	}

	shard.UserData[u.Email()] = u

	return nil
}

func (s *InMemoryStorage) UpdateUser(_ context.Context, u *user.User) error {
	shard, err := s.getShard(u.Email())
	if err != nil {
		return errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.UserData[u.Email()]; !exists {
		return ErrUserNotExists
	}

	shard.UserData[u.Email()] = u

	return nil
}

func (s *InMemoryStorage) GetUser(_ context.Context, email string) (*user.User, error) {
	shard, err := s.getShard(email)
	if err != nil {
		return nil, errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if u, exists := shard.UserData[email]; exists {
		return u, nil
	}

	return nil, ErrUserNotExists
}

func (s *InMemoryStorage) IsAlreadyExistErr(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrUserExists) || errors.Is(err, ErrPlayerExists)
}

func (s *InMemoryStorage) IsNotFoundErr(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrUserNotExists)
}

func (s *InMemoryStorage) SaveGame(_ context.Context, g *game.Game) error {
	shard, err := s.getShard(g.ID().String())
	if err != nil {
		return errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.GameData[g.ID()]; exists {
		return ErrGameExists
	}

	shard.GameData[g.ID()] = g

	return nil
}

func (s *InMemoryStorage) UpdateGame(_ context.Context, g *game.Game) error {
	shard, err := s.getShard(g.ID().String())
	if err != nil {
		return errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.GameData[g.ID()]; !exists {
		return ErrGameNotExists
	}

	shard.GameData[g.ID()] = g

	return nil
}

func (s *InMemoryStorage) GetGame(_ context.Context, id uuid.UUID) (*game.Game, error) {
	shard, err := s.getShard(id.String())
	if err != nil {
		return nil, errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if g, exists := shard.GameData[id]; exists {
		return g, nil
	}

	return nil, ErrGameNotExists
}

func (s *InMemoryStorage) SavePlayer(_ context.Context, p *gmanager.Player) error {
	shard, err := s.getShard(p.UserID())
	if err != nil {
		return errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.PlayerData[p.UserID()]; exists {
		return ErrPlayerExists
	}

	shard.PlayerData[p.UserID()] = p

	return nil
}

func (s *InMemoryStorage) UpdatePlayer(_ context.Context, p *gmanager.Player) error {
	shard, err := s.getShard(p.UserID())
	if err != nil {
		return errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.PlayerData[p.UserID()]; !exists {
		return ErrPlayerNotExists
	}

	shard.PlayerData[p.UserID()] = p

	return nil
}

func (s *InMemoryStorage) GetPlayer(_ context.Context, id string) (*gmanager.Player, error) {
	shard, err := s.getShard(id)
	if err != nil {
		return nil, errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if p, exists := shard.PlayerData[id]; exists {
		return p, nil
	}

	return nil, ErrPlayerNotExists
}
