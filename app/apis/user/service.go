package user

import (
	"context"
	"errors"
	"gewu_jxc/app/tools"
	"gewu_jxc/gen/db"
	"time"
)

func checkSimCodeService(phone string, code string) (db.User, error) {
	var user db.User
	ctx := context.Background()
	realCode := tools.Redis.Get(ctx, "phone:"+phone).Val()
	if realCode != code {
		return user, errors.New("您输入的验证码不正确")
	}

	err := tools.ORM.InsertUser(ctx, db.InsertUserParams{Name: "", Phone: phone, Password: tools.Sha256(tools.RandomCode(99999999))})
	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
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

func regiestSendSimService(phone string) error {
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
