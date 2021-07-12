package account

import (
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
)

func registerEmploy(body registerBody) (Account, error) {
	account := models.Account{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: kit.Sha256(body.Password),
	}

	author := Author{Author: models.Author{ID: 1}}
	employ := Employ{
		Employ:  models.Employ{Boss: true},
		Authors: []Author{author},
	}

	input := Account{
		Account: account,
		Employs: []Employ{employ},
	}
	if err := kit.ORM.Create(&input).Error; err != nil {
		return Account{}, err
	}

	return loadAccount(input.ID)
}
