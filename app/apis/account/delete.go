package account

import (
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func delete(body *deleteBody) error {
	account := models.Account{}
	err := kit.ORM.Where("phone = ?", body.Phone).Take(&account).Error
	if err != nil {
		return errors.New("不存在该手机号用户")
	}

	err = kit.ORM.Where("phone = ?", body.Phone).Delete(&account).Error
	if err != nil {
		return err
	}

	return nil
}
