package user

import (
	"errors"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

func CheckSimCode(c *fiber.Ctx) error {
	var body CheckSimCodeBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	user, err := CheckSimCodeService(&body)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(user))
}

func RegiestSendSim(c *fiber.Ctx) error {
	var body regiestSendSimBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = regiestSendSimService(&body)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(true))
}

func Delete(c *fiber.Ctx) error {
	if !tools.Env.IsDev {
		return errors.New("仅在测试环境才可以使用此API")
	}
	var body deleteBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	err = deleteServer(&body)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(true))
}

func SignInSendSim(c *fiber.Ctx) error {
	var body regiestSendSimBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = signInSendSimService(&body)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(true))
}
