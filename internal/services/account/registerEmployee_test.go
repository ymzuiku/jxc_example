package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/ymzuiku/so"
)

func mockRegisterEmployee(t *testing.T) AccountRes {
	old := mockRegisterCompany(t)
	employeeID := mockLoadCompanysFirstEmployee(t, old.ID)

	val := uuid.NewString()
	data, err := RegisterEmployee(RegisterEmployeeBody{
		CreatorEmployeeID: employeeID,
		Phone:             val,
		Name:              val,
		Password:          val,
	})
	so.Nil(t, err)
	so.True(t, data.ID != 0)
	// var data2 Account
	// orm.DB.Where("id= ?", data.ID).Preload("Employees").Preload("Employees.Authors").Find(&data2)

	employeeID2 := mockLoadCompanysFirstEmployee(t, data.ID)
	permission, err := PermissionLoad(employeeID2)
	so.Nil(t, err)

	so.False(t, permission.Boss)
	so.True(t, permission.CompanyRead)
	so.False(t, permission.EmployeeCreate)
	so.False(t, permission.EmployeeDelete)
	so.False(t, permission.EmployeeRead)
	so.False(t, permission.EmployeeUpdate)

	return data
}

func TestRegisterEmployee(t *testing.T) {
	t.Run("Boss add employee", func(t *testing.T) {
		data := mockRegisterEmployee(t)
		t.Run("Account isn't boss", func(t *testing.T) {
			employeeID := mockLoadCompanysFirstEmployee(t, data.ID)

			val := uuid.NewString()
			_, err := RegisterEmployee(RegisterEmployeeBody{
				CreatorEmployeeID: employeeID,
				Phone:             val,
				Name:              val,
				Password:          val,
			})
			so.NotNil(t, err)
		})
	})
}
func TestAddEmployeeVerify(t *testing.T) {
	t.Run("Boss add employee, if use isn't the company boss, need return error", func(t *testing.T) {
		t.Run("Account isn't boss", func(t *testing.T) {
			// so.Nil(t, fmt.Errorf("ignore_tdd"))
		})
	})
	t.Run("Boss add employee, if set phone is in company's employee", func(t *testing.T) {
		// so.Nil(t, fmt.Errorf("ignore_tdd"))
	})
}

func TestAddEmployeeInput(t *testing.T) {
	t.Run("RegisterEmployeeBody", func(t *testing.T) {
		type Body = RegisterEmployeeBody
		valid := validator.New()
		val := uuid.NewString()
		long := val + val + val

		right := []Body{{CreatorEmployeeID: 10, Phone: "aaaaaaaa", Name: "bbbbbbbbb", Password: "cccccccc"}}

		warn := []Body{
			{},
			{CreatorEmployeeID: 0, Phone: "", Name: "", Password: ""},
			{CreatorEmployeeID: 10, Phone: val, Name: "", Password: ""},
			{CreatorEmployeeID: 10, Phone: val, Name: val, Password: ""},
			{CreatorEmployeeID: 10, Phone: long, Name: val, Password: val},
			{CreatorEmployeeID: 10, Phone: val, Name: long, Password: val},
			{CreatorEmployeeID: 10, Phone: val, Name: val, Password: long},
			{CreatorEmployeeID: 0, Phone: val, Name: val, Password: val},
			{CreatorEmployeeID: 0, Phone: val, Name: val, Password: val},
			{CreatorEmployeeID: 10, Phone: long, Name: val, Password: val},
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
