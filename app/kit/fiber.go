//+build !test

package kit

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var Fiber = fiber.New()

func FiberInit() {
	Fiber.Use(recover.New())

	// Fiber.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Strict",
	// 	Expiration:     1 * time.Hour,
	// 	KeyGenerator:   utils.UUID,
	// }))

	useLogs()
}

func useLogs() {
	format := "[${time}] ${status} - ${latency} ${method} ${path}\n  + ${query}${body}${form:}\n  - ${resBody}\n\n"

	if Env.IsDev {
		Fiber.Use(logger.New(logger.Config{
			Format: format,
		}))
	} else {
		// 按小时分割日志文件
		fileName := fmt.Sprintf("./logs/log_%s.log", time.Now().Format("2006-01-02_15"))

		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil && os.IsNotExist(err) {
			file, err = os.Create(fileName)
		}
		// defer file.Close()

		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		Fiber.Use(logger.New(logger.Config{
			Output: file,
			Format: format,
		}))
	}
}
