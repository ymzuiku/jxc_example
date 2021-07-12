package account

import (
	"fmt"
	"gewu_jxc/app/kit"
	"testing"

	"github.com/stretchr/testify/assert"
)

var phone = "1320000010"
var password = "123123123"
var registerCompanyData = registerCompanyBody{
	Phone:    phone,
	Code:     "999999",
	Company:  "测试企业" + phone,
	Name:     "测试企业用户",
	People:   10,
	Password: password,
}

func TestRegister(t *testing.T) {
	kit.TestInit()

	t.Run("delete account no found", func(t *testing.T) {
		err := remove(removeBody{Phone: "19900000001", Password: password})
		kit.ExitIf(assert.NotNil(t, err))
	})

	t.Run("delete account is found", func(t *testing.T) {
		// 若没有历史账号，假意注册
		if err := registerSendCode(sendCodeBody{Phone: phone}); err != nil {
			fmt.Println(err)
		}

		data, ignoreErr := registerCompany(registerCompanyData)
		hasCompany := true
		if ignoreErr != nil {
			hasCompany = false
			fmt.Printf("假意注册错误，可忽略: %s\n", ignoreErr.Error())
		}

		err := remove(removeBody{Phone: phone, Password: password})

		kit.ExitIf(assert.Nil(t, err))

		if hasCompany {
			err = kit.ORM.Table("account").Where("id = ?", data.ID).Take(nil).Error
			kit.ExitIf(assert.NotNil(t, err))

			err = kit.ORM.Table("company").Where("id = ?", data.Employs[0].CompanyID).Take(nil).Error
			kit.ExitIf(assert.NotNil(t, err))

			err = kit.ORM.Table("employ").Where("account_id = ?", data.ID).Take(nil).Error
			kit.ExitIf(assert.NotNil(t, err))

			err = kit.ORM.Table("employ_author").Where("employ_id = ?", data.Employs[0].ID).Take(nil).Error
			kit.ExitIf(assert.NotNil(t, err))

		}

	})

	t.Run("empty phone register", func(t *testing.T) {
		err := registerSendCode(sendCodeBody{Phone: phone})
		kit.ExitIf(assert.Nil(t, err))

		data, err := registerCompany(registerCompanyData)
		kit.ExitIf(assert.Nil(t, err))

		kit.ExitIf(assert.NotEqual(t, data.Account.ID, 0))
		kit.ExitIf(assert.True(t, len(data.Employs) > 0))
		kit.ExitIf(assert.True(t, len(data.Employs[0].Authors) > 0))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.CompanyRead, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployCreate, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployDelete, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployRead, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployUpdate, true))
	})

}

func TestSignIn(t *testing.T) {
	t.Run("signIn empty with code", func(t *testing.T) {
		err := signInSendCode(sendCodeBody{Phone: phone})
		assert.Nil(t, err)
		_, err = signInWithCode(signInWithCodeBody{
			Phone: "19900000000",
			Code:  "999999",
		})
		kit.ExitIf(assert.NotNil(t, err))
	})

	t.Run("signIn empty with password", func(t *testing.T) {
		_, err := signInWithPassword(signInWithPasswordBody{
			Phone:    "19900000000",
			Password: "123123",
		})
		kit.ExitIf(assert.NotNil(t, err, err))
	})

	t.Run("signIn with code", func(t *testing.T) {
		err := signInSendCode(sendCodeBody{Phone: phone})
		kit.ExitIf(assert.Nil(t, err))

		_, err = signInWithCode(signInWithCodeBody{
			Phone: phone,
			Code:  "999999",
		})
		kit.ExitIf(assert.Nil(t, err))
	})

	t.Run("signIn with password", func(t *testing.T) {
		data, err := signInWithPassword(signInWithPasswordBody{
			Phone:    phone,
			Password: password,
		})

		kit.ExitIf(assert.NotEqual(t, data.Account.ID, 0))
		kit.ExitIf(assert.True(t, len(data.Employs) > 0))
		kit.ExitIf(assert.True(t, len(data.Employs[0].Authors) > 0))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.CompanyRead, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployCreate, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployDelete, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployRead, true))
		kit.ExitIf(assert.Equal(t, data.Employs[0].Permission.EmployUpdate, true))

		kit.ExitIf(assert.Nil(t, err))
	})

}
