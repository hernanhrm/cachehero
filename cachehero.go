package cachehero

import "time"

type UseCase interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	MGet(keys ...string) (map[string]string, error)
}

type Client interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	MGet(keys ...string) ([]interface{}, error)
}
