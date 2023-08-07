package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (any, error)
	Set(string, any, ...int) error
	Forget(string) error
	EmptyByMatch(string) error
	Empty() error
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

type Entry map[string]any

func (c *RedisCache) Has(str string) (bool, error) {
	key := c.makeKey(str)
	conn := c.Conn.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return ok, nil
}

// func (c *RedisCache) Get(string) (any, error)
// func (c *RedisCache) Set(string, any, ...int) error
// func (c *RedisCache) Forget(string) error
// func (c *RedisCache) EmptyByMatch(string) error
// func (c *RedisCache) Empty() error
func (c *RedisCache) makeKey(str string) string {
	return fmt.Sprintf("%s:%s", c.Prefix, str)
}
