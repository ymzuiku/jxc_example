package account

import (
	"context"
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"
)

func signUpSendCode(body *sendCodeBody) error {
	account := models.Account{}
	err := kit.ORM.Where("phone = ?", body.Phone).Take(&account).Error
	if err == nil {
		return errors.New("手机号已注册，请使用该手机号登录")
	}

	code := kit.RandomCode(6)
	kit.Redis.SetEX(context.Background(), "signUp-phone:"+body.Phone, code, time.Minute*10)

	return nil
}
