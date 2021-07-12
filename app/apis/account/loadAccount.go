package account

import (
	"context"
	"encoding/json"
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"time"

	"github.com/mitchellh/mapstructure"
)

const ACCOUNT_CACHE = "accountCache:"

func mergePermission(Autors []models.Author, permission *models.Author) error {
	// 计算权限交集
	var data []map[string]interface{}
	if err := mapstructure.Decode(Autors, &data); err != nil {
		return err
	}

	perm := map[string]interface{}{}
	if err := mapstructure.Decode(permission, &perm); err != nil {
		return err
	}

	// 用于存储所有权限名称
	for _, item := range data {
		for k, v := range item {
			if v == true {
				perm[k] = true
			}
		}
	}

	if err := mapstructure.Decode(perm, &permission); err != nil {
		return err
	}

	return nil
}

func loadAccount(id int32) (Account, error) {
	if cache, err := loadAccountCache(id); err == nil {
		return cache, nil
	}

	var out Account
	if err := kit.ORM.Preload("Employs").Preload("Employs.Company").Preload("Employs.Authors").Where("id = ?", id).Take(&out).Error; err != nil {
		return Account{}, fmt.Errorf("读取账户信息失败: %v\n", err)
	}

	_id := string(id)
	session, err := kit.SessionCreate(_id, _id)
	if err != nil {
		return Account{}, err
	}
	out.Session = session

	// 合并计算每一个 employs 的权限
	for i := 0; i < len(out.Employs); i++ {
		employ := &out.Employs[i]
		var authors []models.Author
		for _, v := range employ.Authors {
			authors = append(authors, v.Author)
		}
		if err := mergePermission(authors, &employ.Permission); err != nil {
			return Account{}, err
		}
	}

	if err := saveAccountCache(id, out); err != nil {
		return Account{}, err
	}

	return out, nil
}

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
