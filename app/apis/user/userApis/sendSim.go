package userApis

import (
	"gewu_jxc/app/apis/user/userServices"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

type _RegisterBody struct {
	Phone string `json:"phone" validate:"required,min=3,max=32"`
}

func SendSim(c *fiber.Ctx) error {
	var body _RegisterBody

	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	err = userServices.SendSim(body.Phone)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(true))
}
