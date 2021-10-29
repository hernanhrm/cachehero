package cachehero

import (
	"fmt"
	"time"
)

type CacheHero struct {
	client Client
}

func New(config Config) (UseCase, error) {
	client, err := newConn(config)
	if err != nil {
		return nil, err
	}

	return CacheHero{client: client}, nil
}

func (c CacheHero) Set(key string, value interface{}, exp time.Duration) error {
	if err := c.client.Set(key, value, exp); err != nil {
		return fmt.Errorf("cachehero: %w", err)
	}

	return nil
}

func (c CacheHero) Get(key string) (interface{}, error) {
	value, err := c.client.Get(key)
	if err != nil {
		return nil, fmt.Errorf("cachehero: %w", err)
	}

	return value, nil
}

func (c CacheHero) Del(key string) error {
	if err := c.client.Del(key); err != nil {
		return fmt.Errorf("cachehero: %w", err)
	}

	return nil
}
