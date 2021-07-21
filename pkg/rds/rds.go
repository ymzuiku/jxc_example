package rds

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func Init() {
	addr := os.Getenv("redisAddr")
	password := os.Getenv("redisPassword")
	dbNum, err := strconv.Atoi(os.Getenv("redisDB"))
	if err != nil {
		log.Fatalln(err)
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNum,
	})
}

type Cache struct {
	prefix string
	redis  *redis.Client
}

var CacheTimeout = time.Hour * 14 * 24

func New(prefix string) Cache {
	return Cache{prefix: prefix, redis: Client}
}

func Del(prefix string, key interface{}) {
	cache := Cache{prefix: prefix, redis: Client}
	cache.Del(key)
}

func (c *Cache) Key(key interface{}) string {
	return fmt.Sprintf("%v::%v", c.prefix, key)
}

func (c *Cache) Get(key interface{}, target interface{}) error {
	data, err := c.redis.GetEx(context.Background(), c.Key(key), CacheTimeout).Bytes()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	return nil
}

func (c *Cache) GetString(key interface{}) string {
	data := c.redis.GetEx(context.Background(), c.Key(key), CacheTimeout).Val()

	return data
}

func (c *Cache) Set(key interface{}, target interface{}) error {
	realKey := c.Key(key)
	switch v := target.(type) {
	case string:
		if err := c.redis.SetEX(context.Background(), realKey, v, CacheTimeout).Err(); err != nil {
			return err
		}
	case int:
		if err := c.redis.SetEX(context.Background(), realKey, v, CacheTimeout).Err(); err != nil {
			return err
		}
	default:
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		if err := c.redis.SetEX(context.Background(), realKey, data, CacheTimeout).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cache) Del(key interface{}) {
	c.redis.Del(context.Background(), c.Key(key))
}

func (c *Cache) Has(key interface{}) error {
	res := c.redis.GetEx(context.Background(), c.Key(key), CacheTimeout)
	err := res.Err()
	if err != nil {
		return err
	}
	val := res.Val()
	if val == "" {
		return fmt.Errorf("session val is error")
	}
	return nil
}

func (c *Cache) Is(key interface{}, val string) bool {
	res := c.redis.GetEx(context.Background(), c.Key(key), CacheTimeout).Val()

	return res == val
}
