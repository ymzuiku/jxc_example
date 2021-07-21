package account

import (
	"errors"

	"github.com/ymzuiku/errox"
)

var errNeedEmployeeID = errors.New("缺少参数 employeeID")

func ChangeComapnyData(employeeID int32) error {
	if employeeID == 0 {
		return errox.Wrap(errNeedEmployeeID)
	}
	return nil
}
