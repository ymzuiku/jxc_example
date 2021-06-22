package userApis

import (
	"errors"
	"gewu_jxc/app/apis/user/userServices"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

type _DeleteBody struct {
	Phone string `json:"phone" validate:"required,min=6,max=32"`
}

func Delete(c *fiber.Ctx) error {
	if !tools.Env.IsDev {
		return errors.New("仅在测试环境才可以使用此API")
	}
	var body _DeleteBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	err = userServices.Delete(body.Phone)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(true))
}
