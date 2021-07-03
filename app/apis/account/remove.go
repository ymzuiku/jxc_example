package account

import (
	"errors"
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"

	"gorm.io/gorm"
)

func remove(body *removeBody) error {
	tx := kit.ORM.Session(&gorm.Session{})
	defer tx.Commit()

	account := models.Account{}
	err := tx.Where("phone = ?", body.Phone).Take(&account).Error
	if err != nil {
		tx.Rollback()
		return errors.New("不存在该手机号用户")
	}

	err = tx.Where("account_id = ?", &account.ID).Delete(&models.Company{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete company: %v", err)
	}

	employ := models.Employ{}

	err = tx.Where("account_id = ?", &account.ID).Delete(&employ).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete company: %v", err)
	}

	err = tx.Where("employ_id = ?", &employ.ID).Delete(&models.EmployActor{}).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete employActor: %v", err)
	}

	err = tx.Where("id = ?", &account.ID).Delete(&account).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete account: %v", err)
	}

	return nil
}
