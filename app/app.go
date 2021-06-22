package app

import (
	"gewu_jxc/app/apis"
	"gewu_jxc/app/tools"
	"log"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	tools.EnvInit()
	tools.PgInit()
	tools.RedisInit()
	tools.Migration(tools.Pg, "sql/migrations")

	tools.Fiber.Use(recover.New())
	tools.UseLogs()
	apis.Init()
	log.Fatal(tools.Fiber.Listen(":3100"))
}
