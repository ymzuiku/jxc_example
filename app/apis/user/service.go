package user

import (
	"context"
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/gen/db"
	"time"
)

func deleteServer(body *deleteBody) error {
	ctx := context.Background()
	_, err := kit.Sql.SelectAccountByPhone(ctx, body.Phone)

	if err != nil {
		return errors.New("不存在该手机号用户")
	}

	err = kit.Sql.DeleteAccountByPhone(ctx, body.Phone)
	if err != nil {
		return err
	}

	return nil
}

func regiestSendSimService(body *regiestSendSimBody) error {
	_, err := kit.Sql.SelectAccountByPhone(context.Background(), body.Phone)

	if err == nil {
		return errors.New("手机号已注册，请使用该手机号登录")
	}

	code := kit.RandomCode(6)

	kit.Redis.SetEX(context.Background(), "regiest-phone:"+body.Phone, code, time.Minute*10)

	return nil
}

func signInSendSimService(body *regiestSendSimBody) error {
	_, err := kit.Sql.SelectAccountByPhone(context.Background(), body.Phone)

	if err != nil {
		return errors.New("手机号未注册，请使用该手机号登录")
	}

	code := kit.RandomCode(6)

	kit.Redis.SetEX(context.Background(), "signIn-phone:"+body.Phone, code, time.Minute*10)

	return nil
}

func CheckSimCodeService(body *CheckSimCodeBody) (db.Account, error) {
	var user db.Account
	ctx := context.Background()
	realCode := kit.Redis.Get(ctx, body.Type+":"+body.Phone).Val()
	if realCode != body.Code {
		return user, errors.New("您输入的验证码不正确")
	}

	err := kit.Sql.InsertAccount(ctx, db.InsertAccountParams{Name: "", Phone: body.Phone, Password: kit.Sha256(kit.RandomCode(8))})
	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}
