package account

import (
	"testing"

	"github.com/ymzuiku/so"
)

func TestLoadPermission(t *testing.T) {
	t.Run("Load new company", func(t *testing.T) {
		data := mockRegisterCompany(t)
		employeeID := mockLoadCompanysFirstEmployee(t, data.ID)
		p, err := PermissionLoad(employeeID)
		so.Nil(t, err)
		// errox.Printf("====== %+v %v\n", p, p.Boss)
		so.True(t, p.Boss)
		so.True(t, p.CompanyRead)
		so.True(t, p.EmployeeCreate)
		so.True(t, p.EmployeeDelete)
		so.True(t, p.EmployeeRead)
		so.True(t, p.EmployeeUpdate)
	})

	t.Run("Load new employee", func(t *testing.T) {
		employee := mockRegisterEmployee(t)
		companys := mockLoadCompanys(t, employee.ID)
		for i := range companys {
			p, err := PermissionLoad(i)
			so.Nil(t, err)
			so.NotEmpty(t, p)
			so.False(t, p.Boss)
			so.True(t, p.CompanyRead)
			so.False(t, p.EmployeeCreate)
			so.False(t, p.EmployeeDelete)
			so.False(t, p.EmployeeRead)
			so.False(t, p.EmployeeUpdate)
		}

	})
}

func TestLoadPermissionVerify(t *testing.T) {
	t.Run("Load empty accountId", func(t *testing.T) {
		permission, err := PermissionLoad(999999988)
		so.NotNil(t, err)
		so.Empty(t, permission)
	})
}
