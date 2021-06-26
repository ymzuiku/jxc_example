package apis

import (
	"gewu_jxc/app/apis/user"
	"gewu_jxc/app/kit"
)

func Init() {
	kit.Fiber.Post("/v1/user/checkSim", user.CheckSimCode)
	kit.Fiber.Post("/v1/user/regiestSendSim", user.RegiestSendSim)
	kit.Fiber.Post("/v1/user/signInSendSim", user.SignInSendSim)
	kit.Fiber.Get("/v0/user/delete", user.Delete)
}
