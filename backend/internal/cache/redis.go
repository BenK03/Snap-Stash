package cache

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// create new redis cache
func NewRedisClient() (*redis.Client, error) {
	addr := strings.TrimSpace(os.Getenv("REDIS_ADDR"))
	if addr == "" {
		addr = "localhost:6379"
	}

	opts := &redis.Options{
		Addr: addr,
	}

	rdb := redis.NewClient(opts)

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return rdb, nil
}

// set bytes
func GetBytes(ctx context.Context, rdb *redis.Client, key string) ([]byte, error) {
	b, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// set bytes with ttl (auto delete after duration)
func SetBytesTTL(ctx context.Context, rdb *redis.Client, key string, value []byte, ttl time.Duration) error {
	return rdb.Set(ctx, key, value, ttl).Err()
}

// deletes
func DelKey(ctx context.Context, rdb *redis.Client, key string) error {
	return rdb.Del(ctx, key).Err()
}
