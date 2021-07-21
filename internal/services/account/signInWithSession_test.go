package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/google/uuid"

	"github.com/ymzuiku/so"
)

func TestSignInSession(t *testing.T) {
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

func TestSignInSessionVerify(t *testing.T) {
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

func TestSignInSessionBody(t *testing.T) {
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
