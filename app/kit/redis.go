package kit

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func RedisInit() {
	addr := os.Getenv("redisAddr")
	password := os.Getenv("redisPassword")
	dbNum, err := strconv.Atoi(os.Getenv("redisDB"))
	if err != nil {
		log.Fatalln(err)
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNum,
	})
}
