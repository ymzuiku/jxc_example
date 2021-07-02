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
	kit.EnvInit(".env")
	kit.PgInit()
	kit.RedisInit()
	kit.Migration(kit.Db)
	kit.FiberInit()
	apis.Init()
	log.Fatal(kit.Fiber.Listen(":3100"))
}
