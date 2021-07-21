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

// const PHONE_IN_60SECOND_ONLYSEND_ONE = "register_phone_in60s"

func RegisterSendCode(body SendCodeBody) error {
	// val := kit.Redis.Get(context.Background(), PHONE_IN_60SECOND_ONLYSEND_ONE+body.Phone).Val()
	// if val == "y" {
	// 	return fmt.Errorf("手机短信发送频繁，请稍后再试")
	// }

	account := models.Account{}
	if err := orm.DB.Where("phone = ?", body.Phone).Take(&account).Error; err == nil {
		return errox.Errorf("手机号已注册，请使用该手机号登录: %w\n", err)
	}

	code := strs.RandomCode(6)
	if err := rds.Client.SetEX(context.Background(), REGISTER_COMPANY_CODE+body.Phone, code, time.Minute).Err(); err != nil {
		return err
	}

	// 设置一个60秒的手机号缓存，确保一个手机号60秒只能发一个此类短信，取消，改为IP限制
	// if err := kit.Redis.SetEX(context.Background(), PHONE_IN_60SECOND_ONLYSEND_ONE+body.Phone, "y", time.Minute*10).Err(); err != nil {
	// 	return err
	// }

	return nil
}
