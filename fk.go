package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"log"
)

type Employ struct {
	models.Employ
	Company models.Company
	Authors []Author `gorm:"many2many:employ_author"`
}

type Author struct {
	models.Author
}

type Account struct {
	models.Account
	Employs []Employ
	Session string `gorm:"-"`
}

func main() {
	kit.TestInit()
	phone := "0006"
	// kit.ORM.Table("account").Where("phone = ?", phone).Update("name", "dog33")

	account := models.Account{
		Name:     "dog",
		Phone:    phone,
		Email:    sql.NullString{String: "mail.qq"},
		Password: "666666",
	}
	company := models.Company{Name: "gewu" + phone, People: 10}
	author := Author{Author: models.Author{ID: 1}}
	employ := Employ{
		Employ:  models.Employ{Boss: true},
		Company: company,
		Authors: []Author{author},
	}

	input := Account{
		Account: account,
		Employs: []Employ{employ},
	}
	if err := kit.ORM.Create(&input).Error; err != nil {
		log.Fatalln(err)
	}

	data, _ := json.Marshal(input)
	fmt.Printf("iiiiiiiiiiiiiiiiiiii %+v", string(data))

	out := Account{}
	if err := kit.ORM.Where("phone = ?", phone).Preload("Employs").Preload("Employs.Company").Preload("Employs.Authors").Take(&out).Error; err != nil {
		log.Fatalln(err)
	}

	data, _ = json.Marshal(out)
	fmt.Printf("ssssssssssssssssssss %+v", string(data))

	if err := kit.ORM.Delete(&input).Error; err != nil {
		log.Fatalln(err)
	}

	ids := make([]int32, 0, len(input.Employs))

	for _, v := range input.Employs {
		if v.Boss {
			ids = append(ids, v.CompanyID)
		}
	}

	if err := kit.ORM.Table("company").Where("id in ?", ids).Delete(nil).Error; err != nil {
		fmt.Println(err)
	}

	// data, _ = json.Marshal(del)
	fmt.Printf("dddddddddddddddddd")
}
