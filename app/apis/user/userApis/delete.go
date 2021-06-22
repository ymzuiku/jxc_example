package userApis

import (
	"context"
	"errors"
	"gewu_jxc/app/tools"

	"github.com/gofiber/fiber/v2"
)

type _DeleteBody struct {
	Phone string `json:"phone" validate:"required,min=6,max=32"`
}

func deleteServer(phone string) error {
	ctx := context.Background()
	_, err := tools.ORM.SelectUserByPhone(ctx, phone)

	if err != nil {
		return errors.New("不存在该手机号用户")
	}

	err = tools.ORM.DeleteUserByPhone(ctx, phone)
	if err != nil {
		return err
	}

	return nil
}

func DeleteApi(c *fiber.Ctx) error {
	if !tools.Env.IsDev {
		return errors.New("仅在测试环境才可以使用此API")
	}
	var body _DeleteBody
	err := c.QueryParser(&body)
	if err != nil {
		return err
	}

	err = deleteServer(body.Phone)

	if err != nil {
		return err
	}

	return c.JSON(tools.Ok(true))
}
