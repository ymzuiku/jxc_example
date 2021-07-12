package account

import "gewu_jxc/models"

type registerBody struct {
	CompanyID int32  `json:"companyID"`
	Phone     string `json:"phone" validate:"required,min=3,max=32"`
	Name      string `json:"name" validate:"required,min=2,max=32"`
	Password  string `json:"password" validate:"required,min=6,max=32"`
}

type registerCompanyBody struct {
	Phone    string `json:"phone" validate:"required,min=3,max=32"`
	Name     string `json:"name" validate:"required,min=2,max=32"`
	Company  string `json:"company" validate:"required,min=2,max=32"`
	People   int32  `json:"people" validate:"required,min=2,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
	Code     string `json:"code" validate:"required, min=6,max=6"`
}

type signInWithCodeBody struct {
	Phone string `json:"phone" validate:"required,min=3,max=32"`
	Code  string `json:"code" validate:"required, min=6,max=6"`
}

type signInWithPasswordBody struct {
	Phone    string `json:"phone" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required, min=6,max=32"`
}

type sendCodeBody struct {
	Phone string `json:"phone" validate:"required,min=3,max=32"`
}

type removeBody struct {
	Phone    string `json:"phone" validate:"required,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type Employ struct {
	models.Employ
	Company    models.Company
	Authors    []Author      `gorm:"many2many:employ_author"`
	Permission models.Author `gorm:"-"`
}

type Author struct {
	models.Author
}

type Account struct {
	models.Account
	Employs []Employ
	Session string `gorm:"-"`
}
