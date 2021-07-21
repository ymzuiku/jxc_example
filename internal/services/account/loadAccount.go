package account

import (
	"github.com/google/uuid"
	"github.com/ymzuiku/errox"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"
)

const ACCOUNT_CACHE = "account"
const SESSION_CACHE = "session"

func LoadAccount(accountID int32) (AccountRes, error) {
	var account AccountRes
	if err := rds.Get(ACCOUNT_CACHE, accountID, &account); err == nil {
		return account, errox.Wrap(err)
	}

	if err := orm.DB.Where("id = ?", accountID).Take(&account).Error; err != nil {
		return AccountRes{}, errox.Errorf("读取账户信息失败: %w\n", err)
	}

	session := uuid.NewString()
	if err := rds.Set(SESSION_CACHE, accountID, session); err != nil {
		return AccountRes{}, errox.Wrap(err)
	}
	account.Session = session

	return account, nil
}

func LoadFullAccount(accountID int32) (Account, error) {
	var account Account

	if err := orm.DB.Preload("Employees").Preload("Employees.Company").Preload("Employees.Authors").Where("id = ?", accountID).Take(&account).Error; err != nil || account.Employees == nil {
		return Account{}, errox.Errorf("读取完整账户信息失败: %w\n", err)
	}

	return account, nil
}
