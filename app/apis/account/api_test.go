package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"testing"
)

func TestSignUp(t *testing.T) {
	kit.EnvInit(".env-test")
	kit.PgInit()
	kit.RedisInit()
	phone := "13200000000"

	t.Run("delete account no found", func(t *testing.T) {
		err := delete(&deleteBody{Phone: "13200000009"})
		if err == nil {
			t.Error("不存在的手机号，未返回正确错误信息")
		}
	})

	t.Run("delete account is found", func(t *testing.T) {
		// 若没有历史账号，假意注册
		_, ignoreErr := signUp(&signUpBody{
			Phone:    phone,
			Code:     "999999",
			Company:  "测试企业",
			Name:     "测试企业用户",
			People:   "10~20",
			Password: "123456",
		})
		if ignoreErr != nil {
			fmt.Printf("假意注册错误，可忽略: %s\n", ignoreErr.Error())
		}

		err := delete(&deleteBody{Phone: phone})
		if err == nil {
			t.Error("存在的手机号，未返回正确错误信息")
		}
	})

	t.Run("signUp sendcode", func(t *testing.T) {
		err := signUpSendCode(&sendCodeBody{Phone: phone})
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("empty phone signUp", func(t *testing.T) {
		_, err := signUp(&signUpBody{
			Phone:    phone,
			Code:     "999999",
			Company:  "测试企业",
			Name:     "测试企业用户",
			People:   "10~20",
			Password: "123456",
		})
		if err != nil {
			t.Error(err.Error())
		}
	})

}
