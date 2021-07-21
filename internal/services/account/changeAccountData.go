package account

import (
	"github.com/ymzuiku/gewu_jxc/internal/models"

	"github.com/ymzuiku/gewu_jxc/pkg/orm"

	"github.com/ymzuiku/errox"
)

func ChangeAccountData(body ChangeAccountDataBody) error {
	if res := orm.DB.Where("id = ?", body.AccountID).Updates(&models.Account{
		Name:  body.Name,
		Email: body.Email,
	}); orm.Error(res) != nil {
		return errox.New("修改账号信息失败")
	}
	return nil
}
