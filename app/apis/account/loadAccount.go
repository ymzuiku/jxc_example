package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func loadAccount(account *models.Account) (accountRes, error) {
	companys := []models.Company{}
	err := kit.ORM.Where("account_id = ?", account.ID).Take(&companys).Error
	if err != nil {
		return accountRes{}, err
	}

	employs := []models.Employ{}
	kit.ORM.Where("account_id = ?", account.ID).Limit(200).Find(&employs)
	if len(employs) == 0 {
		return accountRes{}, fmt.Errorf("find employs = 0")
	}

	var employIDs []int32
	for _, v := range employs {
		employIDs = append(employIDs, v.ID)
	}

	employActors := []models.EmployActor{}
	kit.ORM.Where("employ_id in ?", employIDs).Limit(200).Find(&employActors)
	if len(employs) == 0 {
		return accountRes{}, fmt.Errorf("find actors = 0")
	}

	actorIDs := make([]int32, len(employActors))
	for _, v := range employActors {
		actorIDs = append(actorIDs, v.ID)
	}
	fmt.Printf("bbbbbb %+v\n", employActors)
	fmt.Printf("cccccc %+v\n", actorIDs)

	actors := []models.Actor{}
	kit.ORM.Where("id in ?", actorIDs).Find(&actors)
	if len(actors) == 0 {
		return accountRes{}, fmt.Errorf("actors 未找到")
	}

	// 计算权限交集
	permissions := []map[string]interface{}{}
	kit.ORM.Table("actor_permission").Where("actors_id in ?", actorIDs).Limit(200).Find(&permissions)
	if len(permissions) == 0 {
		return accountRes{}, fmt.Errorf("permissions 未找到")
	}

	permission := models.ActorPermission{
		CompanyRead:  models.OkN,
		EmployCreate: models.OkN,
		EmployDelete: models.OkN,
		EmployRead:   models.OkN,
		EmployUpdate: models.OkN,
	}

	perm := map[string]interface{}{}
	err = mapstructure.Decode(permission, &perm)
	if err != nil {
		return accountRes{}, err
	}

	// 用于存储所有权限名称
	names := []string{}
	for _, item := range permissions {
		names = append(names, item["name"].(string))
		for k, v := range item {
			if v == "y" {
				perm[k] = "y"
			}
		}
	}
	perm["name"] = strings.Join(names, ", ")

	err = mapstructure.Decode(perm, &permission)
	if err != nil {
		return accountRes{}, err
	}

	return accountRes{
		Account:    *account,
		Companys:   companys,
		Employs:    employs,
		Actors:     actors,
		Permission: permission,
	}, nil

}
