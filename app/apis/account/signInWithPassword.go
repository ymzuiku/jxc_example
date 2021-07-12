package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func signInWithPassword(body signInWithPasswordBody) (Account, error) {
	account := models.Account{}
	if res := kit.ORM.Omit("password").Where("phone=? and password = ?", body.Phone, kit.Sha256(body.Password)).Take(&account); res.RowsAffected != 1 {
		return Account{}, fmt.Errorf("您输入的账号或密码不正确")
	}

	return loadAccount(account.ID)
}
