package cache

import (
	"bytes"
	"encoding/gob"
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

func encode(item Entry) ([]byte, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	if err := e.Encode(item); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decode(str string) (Entry, error) {
	b := bytes.Buffer{}
	b.Write([]byte(str))
	d := gob.NewDecoder(&b)
	var item Entry
	if err := d.Decode(&item); err != nil {
		return nil, err
	}

	return item, nil
}

func (c *RedisCache) Get(str string) (any, error) {
	key := c.makeKey(str)
	conn := c.Conn.Get()
	defer conn.Close()

	cacheEntry, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	decoded, err := decode(string(cacheEntry))
	if err != nil {
		return nil, err
	}

	item := decoded[key]

	return item, nil
}
func (c *RedisCache) Set(str string, value any, expires ...int) error {
	key := c.makeKey(str)
	conn := c.Conn.Get()
	defer conn.Close()

	var entry Entry
	entry[key] = value
	encoded, err := encode(entry)
	if err != nil {
		return err
	}

	if len(expires) > 0 {
		_, err := conn.Do("SETEX", key, expires[0], string(encoded))
		if err != nil {
			return err
		}
	} else {
		_, err := conn.Do("SET", key, string(encoded))
		if err != nil {
			return err
		}
	}

	return nil
}
func (c *RedisCache) Forget(str string) error {
	key := c.makeKey(str)
	conn := c.Conn.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}
func (c *RedisCache) EmptyByMatch(str string) error {
	key := c.makeKey(str)
	conn := c.Conn.Get()
	defer conn.Close()

	keys, err := c.getKeys(key)
	if err != nil {
		return err
	}

	for _, x := range keys {
		err := c.Forget(x)
		if err != nil {
			return err
		}
	}

	return nil
}
func (c *RedisCache) Empty() error {
	key := c.makeKey("")
	keys, err := c.getKeys(key)
	if err != nil {
		return err
	}

	for _, x := range keys {
		err := c.Forget(x)
		if err != nil {
			return err
		}
	}

	return nil
}
func (c *RedisCache) makeKey(str string) string {
	return fmt.Sprintf("%s:%s", c.Prefix, str)
}

func (c *RedisCache) getKeys(pattern string) ([]string, error) {
	conn := c.Conn.Get()
	defer conn.Close()

	itr := 0
	keys := []string{}

	for {
		arr, err := redis.Values(conn.Do("SCAN", itr, "MATCH", fmt.Sprintf("%s*", pattern)))
		if err != nil {
			return keys, err
		}

		itr, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if itr == 0 {
			break
		}
	}

	return keys, nil
}
