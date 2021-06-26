package app

import (
	"gewu_jxc/app/apis"
	"gewu_jxc/app/kit"
	"log"
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
	kit.EnvInit()
	kit.PgInit()
	kit.RedisInit()
	kit.Migration(kit.Pg, "gen/migrations")
	kit.FiberInit()
	apis.Init()
	log.Fatal(kit.Fiber.Listen(":3100"))
}
