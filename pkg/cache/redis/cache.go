package redis

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type Cache struct {
	client *redis.Client
	expire time.Duration
}

func (c *Cache) Exists(key string) bool {
	val, err := c.client.Exists(key).Result()
	if err != nil {
		return false
	}
	return val != 0
}

func (c *Cache) Set(key string, value interface{}) error {
	return c.client.Set(key, value, c.expire).Err()
}

func (c *Cache) SetX(key string, value interface{}, expire time.Duration) error {
	return c.client.Set(key, value, expire).Err()
}

func (c *Cache) Get(key string) (interface{}, error) {
	val, err := c.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *Cache) Remove(key string) error {
	return c.client.Del(key).Err()
}

func (c *Cache) Info() string {
	return "No info"
}
