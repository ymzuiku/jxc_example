package account

import (
	"github.com/ymzuiku/gewu_jxc/internal/models"
	"github.com/ymzuiku/gewu_jxc/pkg/orm"
	"github.com/ymzuiku/gewu_jxc/pkg/rds"

	"github.com/mitchellh/mapstructure"
	"github.com/ymzuiku/errox"
)

const PERMISSION_CACHE = "perm"

func PermissionLoad(employeeID int32) (models.Author, error) {
	if employeeID == 0 {
		return models.Author{}, errox.New("请传入正确的ID")
	}

	var permission models.Author
	if err := rds.Get(PERMISSION_CACHE, employeeID, &permission); err == nil {
		return permission, err
	}

	var employee Employee
	if err := orm.DB.Where("id = ?", employeeID).Preload("Authors").Take(&employee).Error; err != nil {
		return models.Author{}, errox.New("读取权限失败")
	}

	data, err := mergePermission(employee.Authors)
	if err != nil {
		return permission, errox.Wrap(err)
	}

	if err := rds.Set(PERMISSION_CACHE, employeeID, data); err != nil {
		return permission, errox.Wrap(err)
	}

	return data, nil

}

func mergePermission(Autors []models.Author) (models.Author, error) {
	// 计算权限交集
	var data []map[string]interface{}
	if err := mapstructure.Decode(Autors, &data); err != nil {
		return models.Author{}, err
	}

	var author models.Author

	perm := map[string]interface{}{}
	if err := mapstructure.Decode(author, &perm); err != nil {
		return author, err
	}

	// 用于存储所有权限名称
	for _, item := range data {
		for k, v := range item {
			if v == true {
				perm[k] = v
			}
		}
	}

	if err := mapstructure.Decode(perm, &author); err != nil {
		return author, err
	}

	return author, nil
}
