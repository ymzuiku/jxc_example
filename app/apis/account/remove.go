package account

import (
	"fmt"
	"gewu_jxc/app/kit"
)

func remove(body removeBody) error {
	// 读取账户信息
	var account Account
	if err := kit.ORM.Where("phone = ? and password = ?", body.Phone, kit.Sha256(body.Password)).Take(&account).Error; err != nil {
		return fmt.Errorf("删除账户关联信息, 不存在该手机号用户: %v", err)
	}

	load, err := loadAccount(account.ID)
	if err != nil {
		return fmt.Errorf("删除账户中，查询账户信息失败: %v\n", err)
	}

	companyIDs := make([]int32, 0, len(load.Employs))
	employIDs := make([]int32, 0, len(load.Employs))
	for _, v := range load.Employs {
		if v.Boss {
			companyIDs = append(companyIDs, v.CompanyID)
		}
		employIDs = append(employIDs, v.ID)
	}

	tx := kit.ORM.Begin()
	if err := tx.Delete(&load).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除账户中，账户删除失败: %v\n", err)
	}

	if err := tx.Delete(&load.Employs).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除账户中，删除 Employ 失败: %+v\n", err)
	}

	if res := tx.Table("company").Where("id in ?", companyIDs).Delete(nil); res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return fmt.Errorf("删除账户中，删除 company 失败: %+v\n", res.Error)
	}

	if res := tx.Table("employ_author").Where("employ_id in ?", employIDs).Delete(nil); res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		return fmt.Errorf("删除账户中，删除 employ_author 失败: %+v\n", res.Error)
	}

	tx.Commit()

	// 清理缓存
	if err := clearAccountCache(account.ID); err != nil {
		return err
	}

	return nil
}
