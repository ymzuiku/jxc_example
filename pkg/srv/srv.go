package srv

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ymzuiku/gewu_jxc/pkg/env"
)

var Fiber = fiber.New()

var _LIMIT_ERROR = func(c *fiber.Ctx) error {
	return fmt.Errorf("当前请求太频繁，请稍后再试")
}

func UseMoneyLimiter() func(*fiber.Ctx) error {
	if !env.IsDev {
		return limiter.New(limiter.Config{
			Max:          1,
			Expiration:   30 * time.Second,
			LimitReached: _LIMIT_ERROR,
		})
	}

	// 开发环境放宽约束
	return limiter.New(limiter.Config{
		Max:          600,
		Expiration:   1 * time.Second,
		LimitReached: _LIMIT_ERROR,
	})

}

func Init() {
	Fiber.Use(recover.New())
	Fiber.Use(compress.New())

	// 生产环境约束每个IP每30秒只能访问 360 次，防止恶意攻击
	if !env.IsDev {
		Fiber.Use(limiter.New(limiter.Config{
			Max:          360,
			Expiration:   30 * time.Second,
			LimitReached: _LIMIT_ERROR,
		}))
	}

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

	if env.IsDev {
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

type ResponseOk struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Ok(data interface{}, msg ...string) ResponseOk {
	if len(msg) > 0 {
		return ResponseOk{Code: 200, Data: data, Msg: msg[0]}
	}
	return ResponseOk{Code: 200, Data: data, Msg: ""}
}
