package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/ymzuiku/so"
)

func mockRegisterCompanyBody(phone string) RegisterCompanyBody {
	return RegisterCompanyBody{
		Phone:    phone,
		Name:     "name_" + phone,
		People:   20,
		Company:  "company_" + phone,
		Password: "123456",
		Code:     "999999",
	}
}

func mockRegisterCompany(t *testing.T) AccountRes {
	phone := uuid.NewString()
	err := RegisterSendCode(SendCodeBody{Phone: phone})
	so.Nil(t, err)
	data, err := RegisterCompany(mockRegisterCompanyBody(phone))
	so.Nil(t, err)
	data.Phone = phone
	return data
}

func TestRegisterAccount(t *testing.T) {
	t.Run("Send Phone Code", func(t *testing.T) {
		phone := uuid.NewString()
		err := RegisterSendCode(SendCodeBody{Phone: phone})
		so.Nil(t, err)
	})

	t.Run("Submit Account Data and Register", func(t *testing.T) {
		phone := uuid.NewString()
		errCode := RegisterSendCode(SendCodeBody{Phone: phone})
		so.Nil(t, errCode)

		acc, err := RegisterCompany(mockRegisterCompanyBody(phone))
		so.Nil(t, err)
		so.NotEmpty(t, acc.ID)
	})

	t.Run("Remove Account", func(t *testing.T) {
		data := mockRegisterCompany(t)

		err := Remove(RemoveBody{Phone: data.Phone, Password: "123456"})
		so.Nil(t, err)

		data2, err := LoadAccount(data.ID)
		so.NotNil(t, err)
		so.Empty(t, data2.ID)
	})
}

func TestRegisterAccountVerify(t *testing.T) {
	t.Run("Submit Account Data and Register, But if Phone is Registered, need return error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		_, err := RegisterCompany(mockRegisterCompanyBody(data.Phone))
		so.NotNil(t, err)
	})

	t.Run("Remove Account empty phone, need return error", func(t *testing.T) {
		err := Remove(RemoveBody{Phone: uuid.NewString(), Password: "123456"})
		so.NotNil(t, err)
	})
	t.Run("Remove Account and password error, need return error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		err := Remove(RemoveBody{Phone: data.Phone, Password: "error123456"})
		so.NotNil(t, err)
	})
}

func TestRegisterInput(t *testing.T) {
	t.Run("RegisterBody", func(t *testing.T) {
		type Body = RegisterCompanyBody
		valid := validator.New()
		val := uuid.NewString()
		long := val + val + val

		right := []Body{{Phone: "aaaaaaaa", Name: "bbbbbbbbb", Password: "cccccccc", Company: "aaaaaa", People: 20, Code: "123456"}}

		warn := []Body{
			{},
			{Phone: "", Name: "", Password: ""},
			{Phone: val, Name: "", Password: ""},
			{Phone: val, Name: val, Password: "", Company: val, Code: "123456"},
			{Phone: long, Name: val, Password: val, Company: val, Code: "123456"},
			{Phone: val, Name: long, Password: val, Company: val, Code: "123456"},
			{Phone: val, Name: val, Password: long, Company: val, Code: "123456"},
			{Phone: val, Name: val, Password: val, Company: long, Code: "123456"},
			{Phone: val, Name: val, Password: val, Company: val, Code: "1234567"},
		}

		for _, v := range right {
			err := valid.Struct(v)
			so.Nil(t, err)
		}
		for _, v := range warn {
			err := valid.Struct(v)
			so.Error(t, err)
		}
	})
}
