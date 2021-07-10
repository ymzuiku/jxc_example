package gewujxcserver

import (
	"fmt"
	"gewu_jxc/app/kit"
	"gewu_jxc/models"
	"testing"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name     string
	age      int
	Companys []models.Company
	Employs  []models.Employ
	Actors   []models.Actor
	Accounts []models.Account
}

func BenchmarkMaps(b *testing.B) {

	var person = Person{
		Name:     "dog",
		age:      20,
		Companys: []models.Company{{ID: 20, Name: "ff", People: models.CpmpanyPeopleLess100, Model: models.CompanyModelFree, DeployModel: models.CompanyDeployModelPrivate, CreatedAt: time.Now(), UpdateAt: time.Now()}},
		Employs:  []models.Employ{{ID: 55555, AccountID: 6666, CompanyID: 6555, Boss: models.OkY, CreatedAt: time.Now(), UpdateAt: time.Now()}},
	}

	var out map[string]interface{}

	for i := 0; i < b.N; i++ {
		if err := mapstructure.Decode(&person, &out); err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkJson(b *testing.B) {

	var person = Person{
		Name:     "dog",
		age:      20,
		Companys: []models.Company{{ID: 20, Name: "ff", People: models.CpmpanyPeopleLess100, Model: models.CompanyModelFree, DeployModel: models.CompanyDeployModelPrivate, CreatedAt: time.Now(), UpdateAt: time.Now()}},
		Employs:  []models.Employ{{ID: 55555, AccountID: 6666, CompanyID: 6555, Boss: models.OkY, CreatedAt: time.Now(), UpdateAt: time.Now()}},
	}

	var out map[string]interface{}

	for i := 0; i < b.N; i++ {
		if err := kit.Parse(&person, &out); err != nil {
			fmt.Println(err)
		}
	}
}
