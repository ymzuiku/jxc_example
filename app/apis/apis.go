package apis

import (
	"gewu_jxc/app/apis/account"
	"gewu_jxc/app/kit"
)

func Init() {
	kit.Fiber.Post("/v1/account/signup/sendcode", account.SignUpSendCodeApi)
	kit.Fiber.Post("/v1/account/signup", account.SignUpApi)

	kit.Fiber.Post("/v1/account/signin/sendcode", account.SignInSendCodeApi)
	kit.Fiber.Post("/v1/account/signin/code", account.SignInWithCodeApi)
	kit.Fiber.Post("/v1/account/signin/password", account.SignInWithPasswordApi)

	if kit.Env.IsDev {
		kit.Fiber.Get("/v0/account/remove", account.RemoveApi)
	}

}
