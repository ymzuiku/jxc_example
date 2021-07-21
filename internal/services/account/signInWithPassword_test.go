package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/google/uuid"

	"github.com/ymzuiku/so"
)

func TestSignInWithPassword(t *testing.T) {
	t.Run("SignIn With Password", func(t *testing.T) {
		data := mockRegisterCompany(t)
		res, err := SignInWithPassword(SignInWithPasswordBody{
			Phone: data.Phone, Password: "123456",
		})
		so.Nil(t, err)
		data2, err := LoadAccount(res.ID)
		so.Nil(t, err)
		so.NotEmpty(t, data2.ID)
	})

}

func TestSignInWithPasswordVerify(t *testing.T) {
	t.Run("SignIn With Password, but phone is error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		_, err := SignInWithPassword(SignInWithPasswordBody{
			Phone: data.Phone + "error", Password: "123456",
		})
		so.NotNil(t, err)
	})
	t.Run("SignIn With Password, but password is error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		_, err := SignInWithPassword(SignInWithPasswordBody{
			Phone: data.Phone, Password: "123456error",
		})
		so.NotNil(t, err)
	})

}

func TestSignInWithPasswordBody(t *testing.T) {

	t.Run("SignInWithPasswordBody", func(t *testing.T) {
		type Body = SignInWithPasswordBody
		valid := validator.New()
		val := uuid.NewString()
		long := val + val

		right := []Body{{Phone: val, Password: val}}

		warn := []Body{
			{},
			{Phone: "", Password: ""},
			{Phone: val},
			{Password: val},
			{Phone: long, Password: val},
			{Phone: val, Password: long},
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
