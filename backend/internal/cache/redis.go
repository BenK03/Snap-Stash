package cache

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Rdb *redis.Client
}

// create new redis cache
func NewRedisCache() (*RedisCache, error) {

	addr := strings.TrimSpace(os.Getenv("REDIS_ADDR"))
	if addr == "" { // if address not set use default
		addr = "localhost:6379"
	}

	// create redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil { // check if properly connected
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	// convert to struct
	cache := RedisCache{
		Rdb: rdb,
	}

	return &cache, nil
}

