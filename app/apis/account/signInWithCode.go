package account

import (
	"context"
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func signInWithCode(body *signInWithCodeBody) (accountRes, error) {
	realCode := kit.Redis.Get(context.Background(), "signIn-phone:"+body.Phone).Val()
	if realCode != body.Code {
		return accountRes{}, errors.New("您输入的账号或验证码不正确")
	}

	account := models.Account{}
	err := kit.ORM.Where("phone = ?", body.Phone).Take(&account).Error
	account.Password = ""
	if err != nil {
		return accountRes{}, err
	}

	return loadAccount(&account)

}
