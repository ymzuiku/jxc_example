package account

import (
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func signInWithPassword(body *signInWithPasswordBody) (models.Account, error) {

	account := models.Account{}
	err := kit.ORM.Where("phone=? and password = ?", body.Phone, body.Password).Take(&account).Error

	if err != nil {
		return account, errors.New("您输入的账号或密码不正确")
	}

	account.Password = ""
	return account, nil
}
