package account

import (
	"context"
	"time"

	"github.com/ymzuiku/gewu_jxc/internal/models"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
	"github.com/ymzuiku/gewu_jxc/pkg/strs"

	"github.com/ymzuiku/errox"
)

func SignInSendCode(body SendCodeBody) error {
	account := models.Account{}
	if err := orm.DB.Where("phone = ?", body.Phone).Take(&account).Error; err != nil {
		return errox.New("手机号未注册，请使用该手机号登录")
	}

	code := strs.RandomCode(6)

	if err := rds.Client.SetEX(context.Background(), "signIn-phone:"+body.Phone, code, time.Minute*10).Err(); err != nil {
		return errox.Wrap(err)
	}

	return nil
}
