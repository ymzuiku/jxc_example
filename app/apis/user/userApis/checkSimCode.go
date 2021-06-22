package userApis

import (
	"gewu_jxc/app/apis/user/userServices"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

type _CheckSimCodeBody struct {
	Phone string `json:"phone" validate:"required,min=3,max=32"`
	Code  string `json:"code" validate:"required, min=4,max=6"`
}

func CheckSimCode(c *fiber.Ctx) error {
	var body _CheckSimCodeBody
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	user, err := userServices.CheckSimCode(body.Phone, body.Code)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(user))
}
