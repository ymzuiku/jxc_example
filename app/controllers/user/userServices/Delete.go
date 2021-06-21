package userServices

import (
	"context"
	"errors"
	"gewu_jxc/app/tools"
)

func Delete(phone string) error {
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
