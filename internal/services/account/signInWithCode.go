package account

import (
	"context"

	"github.com/ymzuiku/gewu_jxc/internal/models"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"

	"github.com/ymzuiku/errox"
)

func SignInWithCode(body SignInWithCodeBody) (AccountRes, error) {
	realCode := rds.Client.Get(context.Background(), "signIn-phone:"+body.Phone).Val()
	if realCode != body.Code {
		return AccountRes{}, errox.New("您输入的账号或验证码不正确")
	}

	account := models.Account{}
	if err := orm.DB.Omit("password").Where("phone = ?", body.Phone).Take(&account).Error; err != nil {
		return AccountRes{}, errox.Wrap(err)
	}

	return LoadAccount(account.ID)

}
