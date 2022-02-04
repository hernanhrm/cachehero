package cachehero

import (
	"context"
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
	if err := r.conn.Set(r.defaultContext, key, value, expiration).Err(); err != nil {
		return fmt.Errorf("redis: could not set key %s, %w", key, err)
	}

	return nil
}

func (r Redis) Get(key string) (string, error) {
	value, err := r.conn.Get(r.defaultContext, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("redis: %w", ErrNotFound)
	}
	if err != nil {
		return "", fmt.Errorf("redis: could not get the key %s, %w", key, err)
	}

	return value, nil
}

func (r Redis) MGet(keys ...string) ([]interface{}, error) {
	value, err := r.conn.MGet(r.defaultContext, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("redis: could not get the keys %s, %w", keys, err)
	}

	return value, nil
}

func (r Redis) Del(key ...string) error {
	err := r.conn.Del(r.defaultContext, key...).Err()
	if errors.Is(err, redis.Nil) {
		return ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("redis: could not delete key %s, %v", key, err)
	}

	return nil
}

func (r Redis) Scan(pattern string, limit int64) ([]string, error) {
	keys, _, err := r.conn.Scan(r.defaultContext, 0, pattern, limit).Result()
	if errors.Is(err, redis.Nil) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("redis: could not get keys with the pattern %s, %v", pattern, err)
	}

	return keys, nil
}
