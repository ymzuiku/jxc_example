package account

import (
	"github.com/ymzuiku/gewu_jxc/internal/models"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"

	"github.com/ymzuiku/errox"
)

// 修改用户权限，传入新的角色ID列表，替换历史
func PermissionChange(body PermissionChangeBody) error {
	if body.EmployeeID == 0 || len(body.AuthorIDs) == 0 {
		return errox.New("请传入正确的ID")
	}
	tx := orm.DB.Begin()

	// 删除所有此员工关联角色
	var employeeAuthor models.EmployeeAuthor
	if err := tx.Where("employee_id = ?", body.EmployeeID).Delete(&employeeAuthor).Error; err != nil {
		return errox.New("创建新权限失败:删除历史权限失败")
	}

	l := len(body.AuthorIDs)
	employeeAuthors := make([]models.EmployeeAuthor, 0, l)
	for i := 0; i < l; i++ {
		employeeAuthors = append(employeeAuthors, models.EmployeeAuthor{EmployeeID: body.EmployeeID, AuthorID: body.AuthorIDs[i]})
	}

	// 添加此员工关联角色
	if res := tx.Create(&employeeAuthors); res.RowsAffected != int64(l) || res.Error != nil {
		return errox.New("创建新权限失败")
	}

	tx.Commit()

	// 清理缓存
	rds.Del(PERMISSION_CACHE, body.EmployeeID)

	return nil
}
