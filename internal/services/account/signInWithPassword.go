package account

import (
	"github.com/ymzuiku/gewu_jxc/internal/models"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/strs"

	"github.com/ymzuiku/errox"
)

func SignInWithPassword(body SignInWithPasswordBody) (AccountRes, error) {
	account := models.Account{}
	if res := orm.DB.Omit("password").Where("phone=? and password = ?", body.Phone, strs.Sha256(body.Password)).Take(&account); orm.Ok(res) != nil {
		return AccountRes{}, errox.Errorf("您输入的账号或密码不正确, %w", res.Error)
	}

	return LoadAccount(account.ID)
}
