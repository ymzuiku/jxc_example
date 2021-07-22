package account

import (
	"errors"

	"github.com/ymzuiku/gewu_jxc/internal/models"

	"github.com/ymzuiku/gewu_jxc/pkg/orm"

	"github.com/ymzuiku/errox"
)

var errChangeAccountDataFail = errors.New("修改账号信息失败")

func ChangeAccountData(body ChangeAccountDataBody) error {
	if err := orm.Ok(orm.DB.Where("id = ?", body.AccountID).Updates(&models.Account{
		Name:  body.Name,
		Email: body.Email,
	})); err != nil {
		return errox.Wrap(errChangeAccountDataFail)
	}
	return nil
}
