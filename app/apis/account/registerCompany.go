package account

import (
	"context"
	"errors"
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"
)

var REGISTER_COMPANY_CODE = "register-phone:"

func registerCompany(body registerCompanyBody) (Account, error) {
	if realCode := kit.Redis.Get(context.Background(), REGISTER_COMPANY_CODE+body.Phone).Val(); realCode != body.Code {
		return Account{}, errors.New("您输入的验证码不正确")
	}

	// 读取管理员角色
	manager := models.Actor{}
	if res := kit.ORM.Where("name = ?", "Manager").Take(&manager); res.RowsAffected != 1 {
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

	if err := tx.Create(&account).Error; err != nil {
		tx.Rollback()
		return Account{}, fmt.Errorf("create account err: %+v", err)
	}

	// 创建企业
	company := models.Company{
		Name:        body.Company,
		People:      body.People,
		Model:       models.CompanyModelFree,
		DeployModel: models.CompanyDeployModelSaas,
	}
	if res := tx.Create(&company); res.RowsAffected != 1 {
		tx.Rollback()
		return Account{}, fmt.Errorf("create company err: %+v", res.Error)
	}

	// 创建员工
	employ := models.Employ{
		AccountID: account.ID,
		CompanyID: company.ID,
		Boss:      models.OkY, // 此人是企业主
	}
	if res := tx.Create(&employ); res.RowsAffected != 1 {
		tx.Rollback()
		return Account{}, fmt.Errorf("create employ err: %+v", res.Error)
	}

	// 创建角色权限映射
	employActor := models.EmployActor{
		EmployID: employ.ID,
		ActorID:  manager.ID,
	}

	if res := tx.Create(&employActor); res.RowsAffected != 1 {
		tx.Rollback()
		return Account{}, fmt.Errorf("create employ_actor err: %+v", res.Error)
	}

	tx.Commit()

	var relAccount models.Account
	if res := kit.ORM.Where("phone = ?", body.Phone).Take(&relAccount); res.RowsAffected != 1 {
		return Account{}, fmt.Errorf("find created account err: %+v", res.Error)
	}

	return loadAccount(relAccount)
}
