package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"

	"github.com/mitchellh/mapstructure"
)

func findPermission(actorIDs []int32, acc *Account) error {
	// 计算权限交集
	var _permissions []models.ActorPermission
	if res := kit.ORM.Table("actor_permission").Where("actor_id in ?", actorIDs).Limit(200).Find(&_permissions); res.RowsAffected == 0 {
		return fmt.Errorf("permissions 未找到")
	}

	var permissions []map[string]interface{}
	if err := mapstructure.Decode(_permissions, &permissions); err != nil {
		return err
	}

	acc.Permission = models.ActorPermission{
		CompanyRead:  models.OkN,
		EmployCreate: models.OkN,
		EmployDelete: models.OkN,
		EmployRead:   models.OkN,
		EmployUpdate: models.OkN,
	}

	perm := map[string]interface{}{}
	if err := mapstructure.Decode(acc.Permission, &perm); err != nil {
		return err
	}

	// 用于存储所有权限名称
	for _, item := range permissions {
		for k, v := range item {
			if v == models.OkY {
				perm[k] = models.OkY
			}
		}
	}

	if err := mapstructure.Decode(perm, &acc.Permission); err != nil {
		return err
	}

	return nil
}

func loadAccount(account models.Account) (Account, error) {
	if cache, err := loadAccountCache(account.ID); err == nil {
		return cache, nil
	}
	fmt.Println(account)
	out := Account{
		Account: account,
	}

	if res := kit.ORM.Where("account_id = ?", out.Account.ID).Limit(200).Find(&out.Employs); res.RowsAffected == 0 {
		return Account{}, errors.New("find employs is empty")
	}

	companyIDs, err := kit.GetIDs(&out.Employs, "CompanyID")
	if err != nil {
		return Account{}, err
	}

	if err := kit.ORM.Where("id in ?", companyIDs).Take(&out.Companys).Error; err != nil {
		return Account{}, err
	}

	employIDs, err := kit.GetIDs(&out.Employs, "ID")
	if err != nil {
		return Account{}, err
	}

	if res := kit.ORM.Where("employ_id in ?", employIDs).Limit(200).Find(&out.EmployActors); res.RowsAffected == 0 {
		return Account{}, fmt.Errorf("find actors is empty")
	}

	actorIDs, err := kit.GetIDs(&out.EmployActors, "ActorID")
	if err != nil {
		return Account{}, err
	}

	if res := kit.ORM.Where("id in ?", actorIDs).Find(&out.Actors); res.RowsAffected == 0 {
		return Account{}, fmt.Errorf("actors 未找到")
	}

	if err := findPermission(actorIDs, &out); err != nil {
		return Account{}, nil
	}

	id := string(out.Account.ID)
	session, err := kit.SessionCreate(id, id)
	if err != nil {
		return Account{}, err
	}
	out.Session = session

	if err := saveAccountCache(out.Account.ID, out); err != nil {
		return Account{}, err
	}

	return out, nil
}

const ACCOUNT_CACHE = "accountCache:"

func loadAccountCache(id int32) (Account, error) {
	data, err := kit.Redis.Get(context.Background(), ACCOUNT_CACHE+string(id)).Bytes()
	if err != nil {
		return Account{}, err
	}

	var account Account
	if err := json.Unmarshal(data, &account); err != nil {
		return account, err
	}

	return account, nil
}

func saveAccountCache(id int32, account Account) error {
	data, err := json.Marshal(account)
	if err != nil {
		return err
	}
	if err := kit.Redis.SetEX(context.Background(), ACCOUNT_CACHE+string(id), data, time.Minute*30).Err(); err != nil {
		return err
	}

	return nil
}

func clearAccountCache(id int32) error {
	kit.Redis.Del(context.Background(), ACCOUNT_CACHE+string(id))
	return nil
}
