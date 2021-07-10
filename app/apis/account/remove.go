package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func remove(body removeBody) error {
	// 读取账户信息
	var account models.Account
	if err := kit.ORM.Where("phone = ?", body.Phone).Take(&account).Error; err != nil {
		return fmt.Errorf("删除账户关联信息, 不存在该手机号用户: %v", err)
	}

	// 查找相关雇员信息
	var employ models.Employ
	if err := kit.ORM.Where("account_id = ?", account.ID).Take(&employ).Error; err != nil {
		return fmt.Errorf("删除过程中，查找雇员失败: %+v", err)
	}

	tx := kit.ORM.Begin()

	// 删除雇员
	if err := tx.Table("employ").Where("account_id = ?", &account.ID).Delete(nil).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除过程中, 删除雇员失败: %+v", err)
	}

	// 删除企业
	if err := tx.Table("company").Where("id = ?", employ.CompanyID).Delete(nil).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除过程中, 删除企业失败: %+v", err)
	}

	// 删除雇员角色中间表
	if err := tx.Table("employ_actor").Where("employ_id = ?", &employ.ID).Delete(nil).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除过程中, 删除雇员角色中间表失败: %+v", err)
	}

	// 删除雇账户
	if err := tx.Table("account").Where("id = ?", &account.ID).Delete(nil).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除过程中, 删除账户: %+v", err)
	}

	tx.Commit()

	// 清理缓存
	if err := clearAccountCache(account.ID); err != nil {
		return err
	}

	return nil
}
