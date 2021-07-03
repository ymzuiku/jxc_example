package account

import (
	"context"
	"errors"
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"
)

func signUp(body *signUpBody) (models.Account, error) {
	realCode := kit.Redis.Get(context.Background(), "signUp-phone:"+body.Phone).Val()
	if realCode != body.Code {
		return models.Account{}, errors.New("您输入的验证码不正确")
	}

	// 读取管理员角色
	manager := models.Actor{}
	res := kit.ORM.Where("name = ?", "Manager").Take(&manager)
	if res.RowsAffected != 1 {
		return models.Account{}, res.Error
	}

	tx := kit.ORM.Begin()

	// 创建账号
	account := models.Account{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: kit.Sha256(body.Password),
		UpdateAt: time.Now(),
	}

	res = tx.Create(&account)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return models.Account{}, fmt.Errorf("create acount err: %+v", res.Error)
	}
	fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaa %+v \n", account)

	// 创建企业
	company := models.Company{
		Name:        body.Company,
		AccountID:   account.ID,
		People:      body.People,
		Model:       models.CompanyModelFree,
		DeployModel: models.CompanyDeployModelSaas,
	}
	res = tx.Create(&company)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return models.Account{}, fmt.Errorf("create company err: %+v", res.Error)
	}

	// 创建员工
	employ := models.Employ{
		AccountID: account.ID,
		CompanyID: company.ID,
	}
	res = tx.Create(&employ)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return models.Account{}, fmt.Errorf("create employ err: %+v", res.Error)
	}

	// 创建角色权限映射
	employActor := models.EmployActor{
		EmployID: employ.ID,
		ActorID:  manager.ID,
	}

	res = tx.Create(&employActor)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return models.Account{}, fmt.Errorf("create company_actor err: %+v", res.Error)
	}

	tx.Commit()

	return account, nil
}
