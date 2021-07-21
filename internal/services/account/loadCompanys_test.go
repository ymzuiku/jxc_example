package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/ymzuiku/so"
)

func mockLoadCompanys(t *testing.T, accountID int32) LoadCompanysRes {
	companys, err := LoadCompanys(accountID)
	so.Nil(t, err)
	so.NotNil(t, companys)
	for _, v := range companys {
		so.NotEmpty(t, v.ID)
		so.NotEmpty(t, v.CreatedAt)
	}

	return companys
}

func mockLoadCompanysFirstEmployee(t *testing.T, accountID int32) int32 {
	companys := mockLoadCompanys(t, accountID)
	var employeeID int32 = 0
	for k := range companys {
		employeeID = k
	}
	return employeeID
}

func TestLoadCompanys(t *testing.T) {
	t.Run("Load boss Comapnys", func(t *testing.T) {
		data := mockRegisterCompany(t)
		companys := mockLoadCompanys(t, data.ID)
		for i := range companys {
			p, err := PermissionLoad(i)
			so.Nil(t, err)
			so.True(t, p.Boss)
		}
	})

	t.Run("Load employee Comapnys", func(t *testing.T) {
		data := mockRegisterEmployee(t)
		companys := mockLoadCompanys(t, data.ID)
		for i := range companys {
			p, err := PermissionLoad(i)
			so.Nil(t, err)
			so.False(t, p.Boss)
		}
	})
}

func TestLoadCompanysBody(t *testing.T) {
	t.Run("LoadCompanysBody", func(t *testing.T) {
		type Body = LoadCompanysBody
		valid := validator.New()

		right := []Body{{AccountID: 10}}

		warn := []Body{
			{},
			{AccountID: 0},
		}

		for _, v := range right {
			so.Nil(t, valid.Struct(v))
		}
		for _, v := range warn {
			so.Error(t, valid.Struct(v))
		}
	})
}
