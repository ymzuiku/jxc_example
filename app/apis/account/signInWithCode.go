package account

import (
	"context"
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func signInWithCode(body *signInWithCodeBody) (models.Account, error) {
	realCode := kit.Redis.Get(context.Background(), "regiest-phone:"+body.Phone).Val()
	if realCode != body.Code {
		return models.Account{}, errors.New("您输入的账号或验证码不正确")
	}

	// account, err := kit.Sql.SelectAccountByPhone(ctx, account.Phone)
	account := models.Account{}
	err := kit.ORM.Where("phone = ?", body.Phone).Take(&account).Error

	if err != nil {
		return account, err
	}

	account.Password = ""
	return account, nil
}
