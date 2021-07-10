package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"
)

func registerEmploy(body registerBody) (Account, error) {
	// 读取管理员角色
	visitor := models.Actor{}
	if res := kit.ORM.Where("name = ?", "Visitor").Take(&visitor); res.RowsAffected != 1 {
		return Account{}, res.Error
	}

	tx := kit.ORM.Begin()

	// 创建账号
	account := models.Account{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: kit.Sha256(body.Password),
		UpdateAt: time.Now(),
	}

	if res := tx.Create(&account); res.RowsAffected != 1 {
		tx.Rollback()
		return Account{}, fmt.Errorf("create acount err: %+v", res.Error)
	}

	// 创建员工
	employ := models.Employ{
		AccountID: account.ID,
		CompanyID: body.CompanyID,
	}
	if res := tx.Create(&employ); res.RowsAffected != 1 {
		tx.Rollback()
		return Account{}, fmt.Errorf("create employ err: %+v", res.Error)
	}

	// 创建角色权限映射
	employActor := models.EmployActor{
		EmployID: employ.ID,
		ActorID:  visitor.ID,
	}

	if res := tx.Create(&employActor); res.RowsAffected != 1 {
		tx.Rollback()
		return Account{}, fmt.Errorf("create company_actor err: %+v", res.Error)
	}

	tx.Commit()

	var relAccount models.Account
	if res := kit.ORM.Where("phone = ?", body.Phone).Take(&relAccount); res.RowsAffected != 1 {
		return Account{}, fmt.Errorf("find created account err: %+v", res.Error)
	}

	return loadAccount(relAccount)
}
