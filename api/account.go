package api

import (
	"github.com/ymzuiku/gewu_jxc/internal/services/account"
	"github.com/ymzuiku/gewu_jxc/pkg/srv"

	"github.com/gofiber/fiber/v2"
)

func accountRegisterSendCode(c *fiber.Ctx) error {
	var body account.SendCodeBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = account.RegisterSendCode(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(true))
}

func accountRegisterEmployee(c *fiber.Ctx) error {
	var body account.RegisterEmployeeBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	account, err := account.RegisterEmployee(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(account))
}

func accountRegisterCompany(c *fiber.Ctx) error {
	var body account.RegisterCompanyBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	account, err := account.RegisterCompany(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(account))
}

// 登录

func accountSignInSendCode(c *fiber.Ctx) error {
	var body account.SendCodeBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = account.SignInSendCode(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(true))
}

func accountSignInWithCode(c *fiber.Ctx) error {
	var body account.SignInWithCodeBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	accountRes, err := account.SignInWithCode(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(accountRes))
}

func accountSignInWithPassword(c *fiber.Ctx) error {
	var body account.SignInWithPasswordBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	accountRes, err := account.SignInWithPassword(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(accountRes))
}

func accountRemove(c *fiber.Ctx) error {
	var body account.RemoveBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	err = account.Remove(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(true))
}

func accountPermissionLoad(c *fiber.Ctx) error {
	var body account.PermissionLoadBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	data, err := account.PermissionLoad(body.EmployeeID)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(data))
}

func accountPermissionChange(c *fiber.Ctx) error {
	var body account.PermissionChangeBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	err = account.PermissionChange(body)

	if err != nil {
		return err
	}

	return c.JSON(srv.Ok(true))
}
