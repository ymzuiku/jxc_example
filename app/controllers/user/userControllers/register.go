package userControllers

import (
	"gewu_jxc/app/controllers/user/userServices"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

func Register() {
	type Body struct {
		Phone string `json:"phone" validate:"required,min=3,max=32"`
	}

	tools.Fiber.Post("/v1/user/register", func(c *fiber.Ctx) error {
		var body Body

		err := c.BodyParser(&body)
		if err != nil {
			return err
		}

		err = userServices.SendSIM(body.Phone)

		if err != nil {
			return err
		}

		return c.JSON(tools.Ok(true))
	})

}
