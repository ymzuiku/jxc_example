package account

import (
	"context"
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"
)

func signInSendCode(body sendCodeBody) error {
	account := models.Account{}
	if err := kit.ORM.Where("phone = ?", body.Phone).Take(&account).Error; err != nil {
		return errors.New("手机号未注册，请使用该手机号登录")
	}

	code := kit.RandomCode(6)

	if err := kit.Redis.SetEX(context.Background(), "signIn-phone:"+body.Phone, code, time.Minute*10).Err(); err != nil {
		return err
	}

	return nil
}
