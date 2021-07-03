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

	var company models.Company
	tx.Where("account_id = ?", &account.ID).Delete(&company)
	if company.ID == 0 {
		tx.Rollback()
		return fmt.Errorf("delete company not found")
	}

	employ := models.Employ{}

	tx.Where("account_id = ?", &account.ID).Delete(&employ)
	if employ.ID == 0 {
		tx.Rollback()
		return fmt.Errorf("delete company not found")
	}

	var employActor models.EmployActor
	tx.Where("employ_id = ?", &employ.ID).Delete(&employActor)
	if employActor.ID == 0 {
		tx.Rollback()
		return fmt.Errorf("delete employActor not found")
	}

	err = tx.Where("id = ?", &account.ID).Delete(&account).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("delete account: %v", err)
	}

	return nil
}
