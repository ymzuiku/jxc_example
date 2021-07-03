package kit

import (
	"math/rand"
	"time"
)

func InitTest() {
	rand.Seed(time.Now().UnixNano())
	EnvInit()
	PgInit()
	RedisInit()
}
