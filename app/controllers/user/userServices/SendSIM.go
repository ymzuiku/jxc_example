package userServices

import (
	"context"
	"errors"
	"fmt"
	"gewu_jxc/app/tools"
	"time"
)

func SendSIM(phone string) error {
	a, err := tools.ORM.SelectUserByPhone(context.Background(), phone)
	fmt.Println(a, err)
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
