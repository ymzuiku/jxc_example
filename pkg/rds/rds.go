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

func Key(prefix string, key interface{}) string {
	return fmt.Sprintf("%v::%v", prefix, key)
}

func Get(prefix string, key interface{}, target interface{}) error {
	data, err := Client.GetEx(context.Background(), Key(prefix, key), CacheTimeout).Bytes()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	return nil
}

func GetString(prefix string, key interface{}) string {
	data := Client.GetEx(context.Background(), Key(prefix, key), CacheTimeout).Val()

	return data
}

func Set(prefix string, key interface{}, target interface{}) error {
	switch v := target.(type) {
	case string:
		if err := Client.SetEX(context.Background(), Key(prefix, key), v, CacheTimeout).Err(); err != nil {
			return err
		}
	case int:
		if err := Client.SetEX(context.Background(), Key(prefix, key), v, CacheTimeout).Err(); err != nil {
			return err
		}
	default:
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		if err := Client.SetEX(context.Background(), Key(prefix, key), data, CacheTimeout).Err(); err != nil {
			return err
		}
	}

	return nil
}

func Del(prefix string, key interface{}) {
	Client.Del(context.Background(), Key(prefix, key))
}

func Has(prefix string, key interface{}) error {
	res := Client.GetEx(context.Background(), Key(prefix, key), CacheTimeout)
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

func Is(prefix string, key interface{}, val string) bool {
	res := Client.GetEx(context.Background(), Key(prefix, key), CacheTimeout).Val()

	return res == val
}
