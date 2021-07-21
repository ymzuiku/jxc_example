package account

import (
	"github.com/ymzuiku/gewu_jxc/internal/models"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/strs"

	"github.com/ymzuiku/errox"
)

func RegisterEmployee(body RegisterEmployeeBody) (AccountRes, error) {
	// kit.ORM.
	permission, err := PermissionLoad(body.CreatorEmployeeID)
	if err != nil {
		return AccountRes{}, errox.Wrap(err)
	}
	if !permission.Boss {
		return AccountRes{}, errox.New("未有创建员工权限")
	}

	var creatorEmployee models.Employee
	if err := orm.DB.Where("id = ?", body.CreatorEmployeeID).Take(&creatorEmployee).Error; err != nil {
		return AccountRes{}, errox.Wrapf(err, "读取企业信息失败")
	}

	input := Account{
		Account: models.Account{
			Name:     body.Name,
			Phone:    body.Phone,
			Password: strs.Sha256(body.Password),
		},
		Employees: []Employee{{
			Employee: models.Employee{CompanyID: creatorEmployee.CompanyID},
			Authors:  []models.Author{{ID: 2}},
		}},
	}
	if err := orm.DB.Create(&input).Error; err != nil {
		return AccountRes{}, errox.Wrap(err)
	}

	return LoadAccount(input.ID)
}
