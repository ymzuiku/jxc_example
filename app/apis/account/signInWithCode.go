package account

import (
	"context"
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func signInWithCode(body signInWithCodeBody) (Account, error) {
	realCode := kit.Redis.Get(context.Background(), "signIn-phone:"+body.Phone).Val()
	if realCode != body.Code {
		return Account{}, errors.New("您输入的账号或验证码不正确")
	}

	account := models.Account{}
	if err := kit.ORM.Omit("password").Where("phone = ?", body.Phone).Take(&account).Error; err != nil {
		return Account{}, err
	}

	return loadAccount(account)

}
