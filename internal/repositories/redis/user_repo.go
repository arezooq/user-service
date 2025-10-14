package redis

import (
	"context"
	"github.com/arezooq/open-utils/db/repository"
	"github.com/redis/go-redis/v9"
)

type UserRedisRepository struct {
	*repository.BaseRedisRepository
}

func NewUserRedisRepository(client *redis.Client, ctx context.Context) *UserRedisRepository {
	return &UserRedisRepository{
		BaseRedisRepository: repository.NewBaseRedisRepository(client, ctx),
	}
}