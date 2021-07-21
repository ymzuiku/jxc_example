package account

import (
	"context"

	"github.com/ymzuiku/gewu_jxc/internal/models"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
	"github.com/ymzuiku/gewu_jxc/pkg/strs"

	"github.com/ymzuiku/errox"
)

const REGISTER_COMPANY_CODE = "register_phone"

func RegisterCompany(body RegisterCompanyBody) (AccountRes, error) {
	if realCode := rds.Client.Get(context.Background(), REGISTER_COMPANY_CODE+body.Phone).Val(); realCode != body.Code {
		return AccountRes{}, errox.New("您输入的验证码不正确")
	}

	input := Account{
		Account: models.Account{
			Name:     body.Name,
			Phone:    body.Phone,
			Password: strs.Sha256(body.Password),
		},
		Employees: []Employee{{
			Employee: models.Employee{},
			Company:  models.Company{Name: body.Company, People: body.People},
			Authors:  []models.Author{{ID: 1}},
		}},
	}

	if res := orm.DB.Create(&input); orm.Error(res) != nil {
		return AccountRes{}, errox.Errorf("创建 account 失败: %w", res.Error)
	}

	return LoadAccount(input.ID)
}
