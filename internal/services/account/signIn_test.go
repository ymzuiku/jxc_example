package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/google/uuid"

	"github.com/ymzuiku/so"
)

func TestSignIn(t *testing.T) {
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
	t.Run("SignIn With Code", func(t *testing.T) {
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
	})
	t.Run("SignIn With Session", func(t *testing.T) {
		data := mockRegisterCompany(t)
		old, err := SignInWithPassword(SignInWithPasswordBody{
			Phone: data.Phone, Password: "123456",
		})
		so.Nil(t, err)

		res, err := SignInWithSession(SignInWithSessionBody{
			AccountID: old.ID, Session: old.Session,
		})
		so.Nil(t, err)

		data2, err := LoadAccount(res.ID)
		so.Nil(t, err)
		so.NotEmpty(t, data2.ID)
		so.NotEmpty(t, data2.Name)
	})
}

func TestSignInVerify(t *testing.T) {
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
	t.Run("SignIn With Session, but id is error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		old, err := SignInWithPassword(SignInWithPasswordBody{
			Phone: data.Phone, Password: "123456",
		})
		so.Nil(t, err)

		_, err = SignInWithSession(SignInWithSessionBody{
			AccountID: old.ID + 10, Session: old.Session,
		})
		so.NotNil(t, err)
	})
	t.Run("SignIn With Session, but session is error", func(t *testing.T) {
		data := mockRegisterCompany(t)
		old, err := SignInWithPassword(SignInWithPasswordBody{
			Phone: data.Phone, Password: "123456",
		})
		so.Nil(t, err)

		_, err = SignInWithSession(SignInWithSessionBody{
			AccountID: old.ID, Session: "aaaaaaaaaaaaaa",
		})
		so.NotNil(t, err)
	})
}

func TestSignInInput(t *testing.T) {
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

	t.Run("SignInWithSessionBody", func(t *testing.T) {
		type Body = SignInWithSessionBody
		valid := validator.New()
		val := uuid.NewString()
		long := val + val

		right := []Body{{AccountID: 12312312, Session: val}}

		warn := []Body{
			{},
			{AccountID: 0, Session: ""},
			{AccountID: 1},
			{AccountID: 0},
			{Session: val},
			{AccountID: 1, Session: long},
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
