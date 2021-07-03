package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"testing"
)

var phone = "13200000004"

func TestSignUp(t *testing.T) {
	kit.InitTest()

	t.Run("delete account no found", func(t *testing.T) {
		err := remove(&removeBody{Phone: "13200000009"})
		if err == nil {
			t.Errorf("不存在的手机号，未返回正确错误信息: %v", err)
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

		err := remove(&removeBody{Phone: phone})
		if err != nil {
			t.Errorf("存在的手机号，未返回正确错误信息: %v", err)
		}
	})

	t.Run("signUp sendcode", func(t *testing.T) {
		err := signUpSendCode(&sendCodeBody{Phone: phone})
		if err != nil {
			t.Error(err)
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
			t.Error(err)
		}
	})

}

func TestSignIn(t *testing.T) {
	t.Run("signIn empty with code", func(t *testing.T) {
		err := signInSendCode(&sendCodeBody{Phone: phone})
		if err != nil {
			t.Error(err)
		}
		_, err = signInWithCode(&signInWithCodeBody{
			Phone: "19900000000",
			Code:  "999999",
		})
		if err == nil {
			t.Error("空账号验证码登录没抛错误")
		}
	})

	t.Run("signIn with code", func(t *testing.T) {
		err := signInSendCode(&sendCodeBody{Phone: phone})
		if err != nil {
			t.Error(err)
		}
		res, err := signInWithCode(&signInWithCodeBody{
			Phone: phone,
			Code:  "999999",
		})
		t.Errorf("%+v", res)
		if err != nil {
			t.Errorf("验证码登录错误: %v", err)
		}
	})

	t.Run("signIn empty with password", func(t *testing.T) {
		_, err := signInWithPassword(&signInWithPasswordBody{
			Phone:    "19900000000",
			Password: "123123",
		})
		if err == nil {
			t.Error("空账号密码登录没抛错误")
		}
	})

}
