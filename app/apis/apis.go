package apis

import (
	"gewu_jxc/app/apis/user/userApis"
	"gewu_jxc/app/tools"
)

func Init() {
	// user
	tools.Fiber.Post("/v1/user/checkSim", userApis.CheckSimCode)
	tools.Fiber.Post("/v1/user/sendSim", userApis.SendSim)
	tools.Fiber.Get("/v0/user/delete", userApis.Delete)
}
