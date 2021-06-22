package apis

import (
	"gewu_jxc/app/apis/user/userApis"
	"gewu_jxc/app/tools"
)

func Init() {
	// user
	tools.Fiber.Post("/v1/user/checkSim", userApis.CheckSimCode)
	tools.Fiber.Post("/v1/user/regiestSendSim", userApis.RegiestSendSim)
	tools.Fiber.Post("/v1/user/signInSendSim", userApis.RegiestSendSim)
	tools.Fiber.Get("/v0/user/delete", userApis.DeleteApi)
}
