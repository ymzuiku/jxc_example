package apis

import (
	"gewu_jxc/app/apis/user"
	"gewu_jxc/app/tools"
)

func Init() {
	tools.Fiber.Post("/v1/user/checkSim", user.CheckSimCode)
	tools.Fiber.Post("/v1/user/regiestSendSim", user.RegiestSendSim)
	tools.Fiber.Post("/v1/user/signInSendSim", user.SignInSendSim)
	tools.Fiber.Get("/v0/user/delete", user.Delete)
}
