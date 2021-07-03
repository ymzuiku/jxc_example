package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func remove(body *removeBody) error {
	var account models.Account
	res := kit.ORM.Where("phone = ?", body.Phone).Take(&account)
	if res.RowsAffected != 1 {
		return fmt.Errorf("不存在该手机号用户")
	}

	tx := kit.ORM.Begin()

	res = tx.Table("company").Where("account_id = ?", &account.ID).Delete(nil)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("delete company err: %+v", res.Error)
	}

	var employ models.Employ
	res = tx.Where("account_id = ?", &account.ID).Take(&employ)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("find employ err: %+v", res.Error)
	}
	res = tx.Table("employ").Where("account_id = ?", &account.ID).Delete(nil)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("delete employ err: %+v", res.Error)
	}

	res = tx.Table("employ_actor").Where("employ_id = ?", &employ.ID).Delete(nil)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("delete employ_actor err: %+v", res.Error)
	}

	res = tx.Table("account").Where("id = ?", &account.ID).Delete(nil)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("delete account err: %+v", res.Error)
	}

	tx.Commit()

	return nil
}
