package app

import (
	"gewu_jxc/app/apis"
	"gewu_jxc/app/tools"
	"log"
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
	tools.EnvInit()
	tools.PgInit()
	tools.RedisInit()
	tools.Migration(tools.Pg, "sql/migrations")
	tools.FiberInit()
	apis.Init()
	log.Fatal(tools.Fiber.Listen(":3100"))
}
