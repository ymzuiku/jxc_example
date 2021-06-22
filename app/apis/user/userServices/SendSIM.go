package userServices

import (
	"context"
	"errors"
	"gewu_jxc/app/tools"
	"time"
)

func SendSIM(phone string) error {
	_, err := tools.ORM.SelectUserByPhone(context.Background(), phone)

	if err == nil {
		return errors.New("手机号已注册，请使用该手机号登录")
	}

	code := "999999"
	if !tools.Env.IsDev {
		code = tools.RandomCode(999999)
	}
	tools.Redis.SetEX(context.Background(), "phone:"+phone, code, time.Minute*10)

	return nil
}
