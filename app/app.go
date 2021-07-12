package app

import (
	"gewu_jxc/app/apis"
	"gewu_jxc/app/kit"
	"log"
	"math/rand"
	"time"

	"github.com/ymzuiku/env_migrate"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
	kit.EnvInit()
	kit.PgInit()
	kit.RedisInit()
	env_migrate.Auto(kit.Db)
	kit.FiberInit()
	apis.Init()
	log.Fatal(kit.Fiber.Listen(":3100"))
}
