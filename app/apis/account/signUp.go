package account

import (
	"context"
	"errors"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"

	"gorm.io/gorm"
)

func signUp(body *signUpBody) (models.Account, error) {
	realCode := kit.Redis.Get(context.Background(), "signUp-phone:"+body.Phone).Val()
	if realCode != body.Code {
		return models.Account{}, errors.New("您输入的验证码不正确")
	}

	// 读取管理员角色
	manager := models.Actor{}
	err := kit.ORM.Where("name = ?", "Manager").Take(&manager).Error
	if err != nil {
		return models.Account{}, err
	}

	tx := kit.ORM.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	// 创建账号
	account := models.Account{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: kit.Sha256(body.Password),
		UpdateAt: time.Now(),
	}

	res := tx.Create(&account)
	if res.RowsAffected != 1 {
		tx.Rollback()
		return models.Account{}, res.Error
	}

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
		return models.Account{}, res.Error
	}

	// 创建员工
	employ := models.Employ{
		AccountID: account.ID,
		CompanyID: company.ID,
	}
	err = tx.Create(&employ).Error
	if err != nil {
		tx.Rollback()
		return models.Account{}, err
	}

	// 创建角色权限映射
	employActor := models.EmployActor{
		EmployID: employ.ID,
		ActorID:  manager.ID,
	}
	err = tx.Create(&employActor).Error

	if err != nil {
		tx.Rollback()
		return models.Account{}, err
	}

	return account, nil
}
