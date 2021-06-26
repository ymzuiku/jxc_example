package user

import (
	"context"
	"errors"
	"gewu_jxc/app/tools"
	"gewu_jxc/gen/db"
	"time"
)

func CheckSimCodeService(body *CheckSimCodeBody) (db.User, error) {
	var user db.User
	ctx := context.Background()
	realCode := tools.Redis.Get(ctx, "phone:"+body.Phone).Val()
	if realCode != body.Code {
		return user, errors.New("您输入的验证码不正确")
	}

	err := tools.ORM.InsertUser(ctx, db.InsertUserParams{Name: "", Phone: body.Phone, Password: tools.Sha256(tools.RandomCode(8))})
	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}

func deleteServer(body *deleteBody) error {
	ctx := context.Background()
	_, err := tools.ORM.SelectUserByPhone(ctx, body.Phone)

	if err != nil {
		return errors.New("不存在该手机号用户")
	}

	err = tools.ORM.DeleteUserByPhone(ctx, body.Phone)
	if err != nil {
		return err
	}

	return nil
}

func regiestSendSimService(body *regiestSendSimBody) error {
	_, err := tools.ORM.SelectUserByPhone(context.Background(), body.Phone)

	if err == nil {
		return errors.New("手机号已注册，请使用该手机号登录")
	}

	code := tools.RandomCode(6)

	tools.Redis.SetEX(context.Background(), "phone:"+body.Phone, code, time.Minute*10)

	return nil
}

func signInSendSimService(body *regiestSendSimBody) error {
	_, err := tools.ORM.SelectUserByPhone(context.Background(), body.Phone)

	if err == nil {
		return errors.New("手机号已注册，请使用该手机号登录")
	}

	code := tools.RandomCode(6)

	tools.Redis.SetEX(context.Background(), "phone:"+body.Phone, code, time.Minute*10)

	return nil
}
