package account

import (
	"database/sql"
	"testing"

	"github.com/ymzuiku/gewu_jxc/internal/models"

	"github.com/ymzuiku/gewu_jxc/pkg/orm"

	"github.com/go-playground/validator"
	"github.com/ymzuiku/so"
)

func TestChangeAccountData(t *testing.T) {
	t.Run("change account name, email", func(t *testing.T) {
		err := ChangeAccountData(ChangeAccountDataBody{
			AccountID: 1,
			Name:      "dog2",
			Email:     sql.NullString{String: "cat2", Valid: true},
		})
		so.Nil(t, err)

		var data models.Account
		err = orm.DB.Where("id = ?", 1).Take(&data).Error
		so.Nil(t, err)
		so.Equal(t, data.Name, "dog2")
		val, err := data.Email.Value()
		so.Nil(t, err)
		so.Equal(t, val, "cat2")
	})
	t.Run("the boss change account people, company", func(t *testing.T) {

	})

}

func TestChangeAccountDataBody(t *testing.T) {
	t.Run("ChangeAccountDataBody", func(t *testing.T) {
		type Body = ChangeAccountDataBody
		valid := validator.New()

		right := []Body{{AccountID: 1, Name: "dog2", Email: sql.NullString{String: "dog@qq.com", Valid: true}}, {AccountID: 1, Name: "dog2", Email: sql.NullString{String: "dog@qq.com", Valid: false}}}

		warn := []Body{
			{},
			{Name: "123"},
			{Email: sql.NullString{String: "dog@qq.com", Valid: true}},
			{Email: sql.NullString{String: "dog@qq.com", Valid: false}},
			{AccountID: 0, Name: "123"},
			{AccountID: 10, Email: sql.NullString{String: "dog@qq.com", Valid: true}},
			{AccountID: 10, Email: sql.NullString{String: "dog@qq.com", Valid: false}},
		}

		for _, v := range right {
			so.Nil(t, valid.Struct(v))
		}
		for _, v := range warn {
			so.Error(t, valid.Struct(v))
		}
	})
}
