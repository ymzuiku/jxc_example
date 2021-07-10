//+build !test

package account

import (
	"gewu_jxc/app/kit"

	"github.com/gofiber/fiber/v2"
)

// 注册

func RegisterSendCodeApi(c *fiber.Ctx) error {
	var body sendCodeBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = registerSendCode(body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(true))
}

func RegisterEmployApi(c *fiber.Ctx) error {
	var body registerBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	account, err := registerEmploy(body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(account))
}

func RegisterCompanyApi(c *fiber.Ctx) error {
	var body registerCompanyBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	account, err := registerCompany(body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(account))
}

// 登录

func SignInSendCodeApi(c *fiber.Ctx) error {
	var body sendCodeBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = signInSendCode(body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(true))
}

func SignInWithCodeApi(c *fiber.Ctx) error {
	var body signInWithCodeBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	accountRes, err := signInWithCode(body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(accountRes))
}

func SignInWithPasswordApi(c *fiber.Ctx) error {
	var body signInWithPasswordBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	accountRes, err := signInWithPassword(body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(accountRes))
}

func RemoveApi(c *fiber.Ctx) error {
	var body removeBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	err = remove(body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(true))
}
