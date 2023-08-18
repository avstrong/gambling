package memory

import (
	"context"
	"hash/fnv"
	"sync"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/user"
	"github.com/google/uuid"
)

var ErrUserExists = errors.New("user already exists")

var ErrUserNotExists = errors.New("user not found")

type memoryShard struct {
	mu   sync.Mutex
	data map[uuid.UUID]*user.User
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
			data: make(map[uuid.UUID]*user.User),
			mu:   sync.Mutex{},
		}
	}

	return &InMemoryStorage{
		shards: shards,
	}
}

func (s *InMemoryStorage) getShard(id uuid.UUID) (*memoryShard, error) {
	hasher := fnv.New32()

	_, err := hasher.Write([]byte(id.String()))
	if err != nil {
		return nil, errors.Wrap(err, "write hash")
	}

	return s.shards[hasher.Sum32()%uint32(len(s.shards))], nil
}

func (s *InMemoryStorage) SaveUser(_ context.Context, u *user.User) error {
	shard, err := s.getShard(u.ID())
	if err != nil {
		return errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if _, exists := shard.data[u.ID()]; exists {
		return ErrUserExists
	}

	shard.data[u.ID()] = u

	return nil
}

func (s *InMemoryStorage) GetUser(_ context.Context, id uuid.UUID) (*user.User, error) {
	shard, err := s.getShard(id)
	if err != nil {
		return nil, errors.Wrap(err, "get shard")
	}

	shard.mu.Lock()
	defer shard.mu.Unlock()

	if u, exists := shard.data[id]; exists {
		return u, nil
	}

	return nil, ErrUserNotExists
}

func (s *InMemoryStorage) IsAlreadyExistErr(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrUserExists)
}

func (s *InMemoryStorage) IsNotFoundErr(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, ErrUserNotExists)
}
