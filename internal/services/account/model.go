package account

import (
	"database/sql"

	"github.com/ymzuiku/gewu_jxc/internal/models"
)

type RegisterEmployeeBody struct {
	CreatorEmployeeID int32  `json:"employeeID" validate:"required,gt=0"`
	Phone             string `json:"phone" validate:"required,min=6,max=36"`
	Name              string `json:"name" validate:"required,min=2,max=36"`
	Password          string `json:"password" validate:"required,min=6,max=36"`
	Email             string `json:"email"`
}

type RegisterCompanyBody struct {
	Phone    string `json:"phone" validate:"required,min=3,max=36"`
	Name     string `json:"name" validate:"required,min=2,max=36"`
	Company  string `json:"company" validate:"required,min=2,max=36"`
	Email    string `json:"email"`
	People   int32  `json:"people" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=36"`
	Code     string `json:"code" validate:"required,min=4,max=6"`
}

type SignInWithCodeBody struct {
	Phone string `json:"phone" validate:"required,min=6,max=36"`
	Code  string `json:"code" validate:"required,min=4,max=6"`
}

type SignInWithPasswordBody struct {
	Phone    string `json:"phone" validate:"required,min=6,max=36"`
	Password string `json:"password" validate:"required,min=6,max=36"`
}

type SignInWithSessionBody struct {
	AccountID int32  `json:"id" validate:"required"`
	Session   string `json:"session" validate:"required,min=6,max=36"`
}

type SendCodeBody struct {
	Phone string `json:"phone" validate:"required,min=6,max=36"`
}

type RemoveBody struct {
	Phone    string `json:"phone" validate:"required,min=6,max=36"`
	Password string `json:"password" validate:"required,min=6,max=36"`
}

type LoadCompanysBody struct {
	AccountID int32 `json:"accountID" validate:"required"`
}

type LoadCompanysRes map[int32]models.Company

type PermissionLoadBody struct {
	EmployeeID int32 `json:"employeeID" validate:"gt=0"`
}
type PermissionChangeBody struct {
	EmployeeID int32   `json:"employeeID" validate:"required"`
	AuthorIDs  []int32 `json:"authorIDs" validate:"required,min=1"`
}

type ChangeAccountDataBody struct {
	AccountID int32          `json:"accountID" validate:"required"`
	Name      string         `json:"name" validate:"required,min=2,max=36"`
	Email     sql.NullString `json:"email" validate:"email"`
}

type AccountRes struct {
	Session string `gorm:"-"`
	models.Account
}

func (a *AccountRes) TableName() string {
	return "account"
}

type Account struct {
	models.Account
	Employees []Employee
	Session   string `gorm:"-"`
}

type Employee struct {
	models.Employee
	Company models.Company
	Authors []models.Author `gorm:"many2many:employee_author"`
}
