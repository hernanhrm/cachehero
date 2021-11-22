package cachehero

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	if !isPrimitiveType(value) {
		valueBytes, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf("redis: could not prepare value to set with key %s, %w", key, err)
		}

		value = valueBytes
	}

	if err := c.client.Set(key, value, exp); err != nil {
		return fmt.Errorf("cachehero: %w", err)
	}

	return nil
}

func (c CacheHero) Get(key string) (string, error) {
	value, err := c.client.Get(key)
	if err != nil {
		return "", fmt.Errorf("cachehero: %w", err)
	}

	return value, nil
}

func (c CacheHero) MGet(keys ...string) (map[string]string, error) {
	valuesMap := make(map[string]string)
	if len(keys) <= 0 {
		return valuesMap, ErrNotEnteredKeys
	}

	values, err := c.client.MGet(keys...)
	if err != nil {
		return nil, fmt.Errorf("cachehero: %w", err)
	}

	for i := range values {
		if values[i] == nil {
			valuesMap[keys[i]] = ""
			continue
		}
		valuesMap[keys[i]] = values[i].(string)
	}

	return valuesMap, nil
}

func (c CacheHero) Del(key string) error {
	if err := c.client.Del(key); err != nil {
		return fmt.Errorf("cachehero: %w", err)
	}

	return nil
}

func isPrimitiveType(value interface{}) bool {
	reflectedValueKind := reflect.TypeOf(value).Kind()
	return reflectedValueKind != reflect.Struct && reflectedValueKind != reflect.Map
}
