package userApis

import (
	"gewu_jxc/app/apis/user/userServices"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

type _SignInSendSimBody struct {
	Phone string `json:"phone" validate:"required,min=3,max=32"`
}

func SignInSendSim(c *fiber.Ctx) error {
	var body _SignInSendSimBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = userServices.RegiestSendSim(body.Phone)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(true))
}
