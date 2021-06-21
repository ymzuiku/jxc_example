package userServices

import (
	"context"
	"errors"
	"gewu_jxc/app/tools"
	"gewu_jxc/sql/db"
)

func CheckSimCode(phone string, code string) (db.User, error) {
	var user db.User
	ctx := context.Background()
	realCode := tools.Redis.Get(ctx, "phone:"+phone).Val()
	if realCode != code {
		return user, errors.New("您输入的验证码不正确")
	}

	err := tools.ORM.InsertUser(ctx, db.InsertUserParams{Name: "", Phone: phone, Password: tools.Sha256(tools.RandomCode(99999999))})
	if err != nil {
		return user, err
	}

	user.Password = ""
	return user, nil
}
