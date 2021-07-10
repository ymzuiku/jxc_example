package account

import "gewu_jxc/models"

type registerBody struct {
	CompanyID int32  `json:"companyID"`
	Phone     string `json:"phone" validate:"required,min=3,max=32"`
	Name      string `json:"name" validate:"required,min=2,max=32"`
	Password  string `json:"password" validate:"required,min=6,max=32"`
}

type registerCompanyBody struct {
	Phone    string               `json:"phone" validate:"required,min=3,max=32"`
	Name     string               `json:"name" validate:"required,min=2,max=32"`
	Company  string               `json:"company" validate:"required,min=2,max=32"`
	People   models.CpmpanyPeople `json:"people" validate:"required,min=2,max=32"`
	Password string               `json:"password" validate:"required,min=6,max=32"`
	Code     string               `json:"code" validate:"required, min=6,max=6"`
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
	Phone string `json:"phone" validate:"required,min=6,max=32"`
}

type Account struct {
	Session      string
	Account      models.Account
	Employs      []models.Employ        `json:"employs"`
	Companys     []models.Company       `json:"companys"`
	EmployActors []models.EmployActor   `json:"employActors"`
	Actors       []models.Actor         `json:"actors"`
	Permission   models.ActorPermission `json:"permission"`
}
