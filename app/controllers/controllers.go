package controllers

import (
	"gewu_jxc/app/controllers/user"
	"gewu_jxc/app/tools"
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(port string) {
	tools.Fiber.Use(recover.New())
	if tools.Env.IsDev {
		tools.Fiber.Use(logger.New(logger.Config{}))
	}
	user.UserInit()
	log.Fatal(tools.Fiber.Listen(port))
}
