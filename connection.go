package cachehero

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

const redisDriver = "redis"

func newConn(config Config) (Client, error) {
	switch config.Driver {
	case redisDriver:
		return newRedisConn(config)
	default:
		return nil, fmt.Errorf("conn: driver not implemented")
	}
}

func newRedisConn(config Config) (*Redis, error) {
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	db, err := strconv.Atoi(config.Database)
	if err != nil {
		return nil, fmt.Errorf("conn: could not parse redis database %d, %w", config.Database, err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: config.Username,
		Password: config.Password,
		DB:       db,
	})

	return newRedis(redisClient), nil
}
