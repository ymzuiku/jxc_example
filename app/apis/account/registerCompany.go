package account

import (
	"context"
	"errors"
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

var REGISTER_COMPANY_CODE = "register-phone:"

func registerCompany(body registerCompanyBody) (Account, error) {
	if realCode := kit.Redis.Get(context.Background(), REGISTER_COMPANY_CODE+body.Phone).Val(); realCode != body.Code {
		return Account{}, errors.New("您输入的验证码不正确")
	}

	company := models.Company{Name: body.Company, People: body.People}
	author := Author{Author: models.Author{ID: 1}}
	employ := Employ{
		Employ:  models.Employ{Boss: true},
		Company: company,
		Authors: []Author{author},
	}

	input := Account{
		Account: models.Account{
			Name:     body.Name,
			Phone:    body.Phone,
			Password: kit.Sha256(body.Password),
		},
		Employs: []Employ{employ},
	}
	if res := kit.ORM.Create(&input); res.Error != nil || res.RowsAffected != 1 {
		return Account{}, fmt.Errorf("创建 account 失败: %v", res.Error)
	}

	return loadAccount(input.ID)
}
