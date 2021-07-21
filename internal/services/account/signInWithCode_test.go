package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/google/uuid"

	"github.com/ymzuiku/so"
)

func TestSignInWithCode(t *testing.T) {
	data := mockRegisterCompany(t)
	err := SignInSendCode(SendCodeBody{Phone: data.Phone})
	so.Nil(t, err)

	res, err := SignInWithCode(SignInWithCodeBody{
		Phone: data.Phone, Code: "999999",
	})
	so.Nil(t, err)

	data2, err := LoadAccount(res.ID)
	so.Nil(t, err)
	so.NotEmpty(t, data2.ID)

}

func TestSignInCodeVerify(t *testing.T) {
	t.Run("SignIn With Code, but phone is error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		err := SignInSendCode(SendCodeBody{Phone: data.Phone})
		so.Nil(t, err)

		_, err = SignInWithCode(SignInWithCodeBody{
			Phone: data.Phone + "2", Code: "999999",
		})

		so.NotNil(t, err)

	})
	t.Run("SignIn With Code, but code is error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		err := SignInSendCode(SendCodeBody{Phone: data.Phone})
		so.Nil(t, err)

		_, err = SignInWithCode(SignInWithCodeBody{
			Phone: data.Phone, Code: "999991",
		})

		so.NotNil(t, err)
	})
}

func TestSignInCodeBody(t *testing.T) {
	t.Run("SignInWithCodeBody", func(t *testing.T) {
		type Body = SignInWithCodeBody
		valid := validator.New()
		val := uuid.NewString()
		long := val + val + val

		right := []Body{{Phone: val, Code: "123456"}}

		warn := []Body{
			{},
			{Phone: "", Code: ""},
			{Phone: val, Code: ""},
			{Phone: long, Code: "123123"},
			{Phone: val, Code: long},
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
