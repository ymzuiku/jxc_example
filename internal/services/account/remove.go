package account

import (
	"github.com/ymzuiku/errox"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
	"github.com/ymzuiku/gewu_jxc/pkg/strs"
)

func Remove(body RemoveBody) error {
	// 读取账户信息
	var account Account
	if err := orm.DB.Where("phone = ? and password = ?", body.Phone, strs.Sha256(body.Password)).Take(&account).Error; err != nil {
		return errox.Errorf("删除账户关联信息, 不存在该手机号用户: %w", err)
	}

	load, err := LoadFullAccount(account.ID)
	if err != nil {
		return errox.Errorf("删除账户中，查询账户信息失败: %w\n", err)
	}

	companyIDs := make([]int32, 0, len(load.Employees))
	employeeIDs := make([]int32, 0, len(load.Employees))
	for _, v := range load.Employees {

		if v.Authors != nil && v.Authors[0].Boss {
			companyIDs = append(companyIDs, v.CompanyID)
		}
		employeeIDs = append(employeeIDs, v.ID)
	}

	tx := orm.DB.Begin()
	if err := tx.Delete(&load).Error; err != nil {
		tx.Rollback()
		return errox.Errorf("删除账户中，删除 Account 失败: %w\n", err)
	}

	if err := tx.Delete(&load.Employees).Error; err != nil {
		tx.Rollback()
		return errox.Errorf("删除账户中，删除 Employee 失败: %w\n", err)
	}

	if res := tx.Table("company").Where("id in ?", companyIDs).Delete(nil); orm.Error(res) != nil {
		tx.Rollback()
		return errox.Errorf("删除账户中，删除 Company 失败: %w\n", res.Error)
	}

	if res := tx.Table("employee_author").Where("employee_id in ?", employeeIDs).Delete(nil); orm.Error(res) != nil {
		tx.Rollback()
		return errox.Errorf("删除账户中，删除 Employee_author 失败: %w\n", res.Error)
	}

	tx.Commit()

	// 清理缓存
	rds.Del(ACCOUNT_CACHE, account.ID)

	return nil
}
