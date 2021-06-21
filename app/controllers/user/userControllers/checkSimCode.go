package userControllers

import (
	"gewu_jxc/app/controllers/user/userServices"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

func CheckSimCode() {
	type Body struct {
		Phone string `json:"phone" validate:"required,min=3,max=32"`
		Code  string `json:"code" validate:"required, min=4,max=6"`
	}

	tools.Fiber.Post("/v1/user/checkSim", func(c *fiber.Ctx) error {
		var body Body
		err := c.BodyParser(&body)
		if err != nil {
			return err
		}

		user, err := userServices.CheckSimCode(body.Phone, body.Code)

		if err != nil {
			return err
		}

		return c.JSON(tools.Ok(user))
	})
}
