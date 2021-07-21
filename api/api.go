package api

import (
	"github.com/ymzuiku/gewu_jxc/pkg/env"
	"github.com/ymzuiku/gewu_jxc/pkg/srv"
)

func Init() {
	srv.Fiber.Post("/v1/account/register/sendcode", srv.UseMoneyLimiter(), accountRegisterSendCode)
	srv.Fiber.Post("/v1/account/register/company", accountRegisterCompany)
	srv.Fiber.Post("/v1/account/register/employee", accountRegisterEmployee)
	srv.Fiber.Post("/v1/account/signin/sendcode", srv.UseMoneyLimiter(), accountSignInSendCode)
	srv.Fiber.Post("/v1/account/signin/code", accountSignInWithCode)
	srv.Fiber.Post("/v1/account/signin/password", accountSignInWithPassword)

	srv.Fiber.Post("/v1/account/permission/load", accountPermissionLoad)
	srv.Fiber.Post("/v1/account/permission/change", accountPermissionChange)

	if env.IsDev {
		srv.Fiber.Get("/v0/account/remove", accountRemove)
	}

}
