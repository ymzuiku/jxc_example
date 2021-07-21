package account

import (
	"github.com/ymzuiku/errox"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
)

const COMPANYS = "companys"

func LoadCompanys(accountID int32) (LoadCompanysRes, error) {
	if accountID == 0 {
		return LoadCompanysRes{}, errox.New("请传入正确的ID")
	}
	cache := rds.New(COMPANYS)

	companysRes := LoadCompanysRes{}

	if err := cache.Get(accountID, &companysRes); err == nil {
		return companysRes, nil
	}

	var account Account
	if err := orm.DB.Where("id", accountID).Preload("Employees").Preload("Employees.Company").Find(&account).Error; err != nil {
		return companysRes, errox.Wrap(err)
	}

	for _, v := range account.Employees {
		companysRes[v.ID] = v.Company
	}

	if err := cache.Set(accountID, companysRes); err != nil {
		return companysRes, nil
	}

	return companysRes, nil
}
