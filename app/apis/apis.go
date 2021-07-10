package apis

import (
	"gewu_jxc/app/apis/account"
	"gewu_jxc/app/kit"
)

func Init() {
	kit.Fiber.Post("/v1/account/register/sendcode", account.RegisterSendCodeApi)
	kit.Fiber.Post("/v1/account/register/company", account.RegisterCompanyApi)
	kit.Fiber.Post("/v1/account/register/employ", account.RegisterEmployApi)

	kit.Fiber.Post("/v1/account/signin/sendcode", account.SignInSendCodeApi)
	kit.Fiber.Post("/v1/account/signin/code", account.SignInWithCodeApi)
	kit.Fiber.Post("/v1/account/signin/password", account.SignInWithPasswordApi)

	if kit.Env.IsDev {
		kit.Fiber.Get("/v0/account/remove", account.RemoveApi)
	}

}
