package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/ymzuiku/so"
)

func TestPermissionChange(t *testing.T) {
	t.Run("change master to visiter", func(t *testing.T) {
		data := mockRegisterCompany(t)
		employeeID := mockLoadCompanysFirstEmployee(t, data.ID)
		p, err := PermissionLoad(employeeID)
		so.Nil(t, err)
		so.True(t, p.Boss)
		so.True(t, p.CompanyRead)
		so.True(t, p.EmployeeCreate)
		so.True(t, p.EmployeeDelete)
		so.True(t, p.EmployeeRead)
		so.True(t, p.EmployeeUpdate)

		err = PermissionChange(PermissionChangeBody{
			EmployeeID: employeeID,
			AuthorIDs:  []int32{2},
		})
		so.Nil(t, err)

		p2, err2 := PermissionLoad(employeeID)
		so.Nil(t, err2)
		so.False(t, p2.Boss)
		so.True(t, p2.CompanyRead)
		so.False(t, p2.EmployeeCreate)
		so.False(t, p2.EmployeeDelete)
		so.False(t, p2.EmployeeRead)
		so.False(t, p2.EmployeeUpdate)
	})

	t.Run("change vister to master", func(t *testing.T) {
		data := mockRegisterCompany(t)
		employeeID := mockLoadCompanysFirstEmployee(t, data.ID)
		p, err := PermissionLoad(employeeID)
		so.Nil(t, err)
		so.True(t, p.Boss)
		so.True(t, p.CompanyRead)
		so.True(t, p.EmployeeCreate)
		so.True(t, p.EmployeeDelete)
		so.True(t, p.EmployeeRead)
		so.True(t, p.EmployeeUpdate)

		err = PermissionChange(PermissionChangeBody{
			EmployeeID: employeeID,
			AuthorIDs:  []int32{1},
		})
		so.Nil(t, err)

		p2, err2 := PermissionLoad(employeeID)
		so.Nil(t, err2)
		so.True(t, p2.Boss)
		so.True(t, p2.CompanyRead)
		so.True(t, p2.EmployeeCreate)
		so.True(t, p2.EmployeeDelete)
		so.True(t, p2.EmployeeRead)
		so.True(t, p2.EmployeeUpdate)
	})
	t.Run("change vister and master", func(t *testing.T) {
		data := mockRegisterCompany(t)
		employeeID := mockLoadCompanysFirstEmployee(t, data.ID)
		p, err := PermissionLoad(employeeID)
		so.Nil(t, err)
		so.True(t, p.Boss)
		so.True(t, p.CompanyRead)
		so.True(t, p.EmployeeCreate)
		so.True(t, p.EmployeeDelete)
		so.True(t, p.EmployeeRead)
		so.True(t, p.EmployeeUpdate)

		err = PermissionChange(PermissionChangeBody{
			EmployeeID: employeeID,
			AuthorIDs:  []int32{1, 2},
		})
		so.Nil(t, err)

		p2, err2 := PermissionLoad(employeeID)
		so.Nil(t, err2)
		so.True(t, p2.Boss)
		so.True(t, p2.CompanyRead)
		so.True(t, p2.EmployeeCreate)
		so.True(t, p2.EmployeeDelete)
		so.True(t, p2.EmployeeRead)
		so.True(t, p2.EmployeeUpdate)
	})
}

func TestPermissionChangeBody(t *testing.T) {
	type Body = PermissionChangeBody

	valid := validator.New()
	right := []Body{{EmployeeID: 1, AuthorIDs: []int32{1, 2, 3}}}

	warn := []Body{
		{},
		{EmployeeID: 0},
		{EmployeeID: 10},
		{AuthorIDs: []int32{}},
		{AuthorIDs: []int32{}},
		{AuthorIDs: []int32{1}},
		{AuthorIDs: []int32{1, 2}},
	}

	for _, v := range right {
		so.Nil(t, valid.Struct(v))
	}
	for _, v := range warn {
		so.Error(t, valid.Struct(v))
	}
}
