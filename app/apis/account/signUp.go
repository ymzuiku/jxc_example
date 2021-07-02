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
	realCode := kit.Redis.Get(context.Background(), "regiest-phone:"+body.Phone).Val()
	if realCode != body.Code {
		return models.Account{}, errors.New("您输入的验证码不正确")
	}

	tx := kit.ORM.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	// manager, err := kit.Sql.SelectActorByName(ctx, "Manager")
	manager := models.Actor{}
	err := tx.Where("name = ?", "Manager").Take(&manager).Error

	if err != nil {
		tx.Rollback()
		return models.Account{}, err
	}

	account := models.Account{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: kit.Sha256(body.Password),
		UpdateAt: time.Now(),
	}

	res := tx.Create(&account)
	if res.Error != nil {
		tx.Rollback()
		return models.Account{}, res.Error
	}

	company := models.Company{
		Name:        body.Company,
		AccountID:   account.ID,
		People:      body.People,
		Model:       models.CompanyModelFree,
		DeployModel: models.CompanyDeployModelSaas,
	}
	err = tx.Omit("name, account_id, people").Create(&company).Error
	if err != nil {
		tx.Rollback()
		return models.Account{}, err
	}

	employ := models.Employ{
		AccountID: account.ID,
		CompanyID: company.ID,
	}
	err = tx.Omit("account_id, company_id").Create(&employ).Error
	if err != nil {
		tx.Rollback()
		return models.Account{}, err
	}

	employActor := models.EmployActor{
		EmployID: employ.ID,
		ActorID:  manager.ID,
	}
	err = tx.Omit("employ_id, actor_id").Create(&employActor).Error

	if err != nil {
		tx.Rollback()
		return models.Account{}, err
	}

	return account, nil
}
