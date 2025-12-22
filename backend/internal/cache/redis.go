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

func NewRedisCache() (*RedisCache, error) {
	addr := strings.TrimSpace(os.Getenv("REDIS_ADDR"))
	if addr == "" { // if address not set use default
		addr = "localhost:6379"
	}

	// create redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})


}
