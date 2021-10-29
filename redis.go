package cachehero

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	conn           *redis.Client
	defaultContext context.Context
}

func newRedis(conn *redis.Client) *Redis {
	return &Redis{conn: conn, defaultContext: context.Background()}
}

func (r Redis) Set(key string, value interface{}, expiration time.Duration) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("redis: could not prepare value to set with key %s, %w", key, err)
	}

	if err := r.conn.Set(r.defaultContext, key, valueBytes, expiration).Err(); err != nil {
		return fmt.Errorf("redis: could not set key %s, %w", key, err)
	}

	return nil
}

func (r Redis) Get(key string) (interface{}, error) {
	value, err := r.conn.Get(r.defaultContext, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("redis: %w", ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("redis: could not get the key %s, %w", key, err)
	}

	return value, nil
}

func (r Redis) Del(key string) error {
	if err := r.conn.Del(r.defaultContext, key); err != nil {
		return fmt.Errorf("redis: could not delete key %s, %w", key, err)
	}

	return nil
}
