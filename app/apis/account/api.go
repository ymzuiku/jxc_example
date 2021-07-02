package account

import (
	"gewu_jxc/app/kit"

	"github.com/gofiber/fiber/v2"
)

// 注册

func SignUpSendCodeApi(c *fiber.Ctx) error {
	var body sendCodeBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = signUpSendCode(&body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(true))
}

func SignUpApi(c *fiber.Ctx) error {
	var body signUpBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	account, err := signUp(&body)

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

	err = signInSendCode(&body)

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

	account, err := signInWithCode(&body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(account))
}

func SignInWithPasswordApi(c *fiber.Ctx) error {
	var body signInWithPasswordBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	account, err := signInWithPassword(&body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(account))
}

func DeleteApi(c *fiber.Ctx) error {
	var body deleteBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	err = delete(&body)

	if err != nil {
		return err
	}

	return c.JSON(kit.Ok(true))
}
