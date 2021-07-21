package account

import (
	"errors"

	"github.com/ymzuiku/errox"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
)

const COMPANYS = "companys"

var errNeedAccountID = errors.New("请传入正确的ID")

func LoadCompanys(accountID int32) (LoadCompanysRes, error) {
	if accountID == 0 {
		return LoadCompanysRes{}, errox.Wrap(errNeedAccountID)
	}

	companysRes := LoadCompanysRes{}

	if err := rds.Get(COMPANYS, accountID, &companysRes); err == nil {
		return companysRes, nil
	}

	var account Account
	if err := orm.DB.Where("id", accountID).Preload("Employees").Preload("Employees.Company").Find(&account).Error; err != nil {
		return companysRes, errox.Wrap(err)
	}

	for _, v := range account.Employees {
		companysRes[v.ID] = v.Company
	}

	if err := rds.Set(COMPANYS, accountID, companysRes); err != nil {
		return companysRes, nil
	}

	return companysRes, nil
}
