package fixtures

import (
	"time"

	"github.com/swimlibre/backend-api/model"
	"gorm.io/gorm"
)

func LoadFixtures(db *gorm.DB) {
	RunMigrations(db)
	LoadData(db)
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Organization{})
	db.AutoMigrate(&model.Season{})
	db.AutoMigrate(&model.Team{})
	db.AutoMigrate(&model.Person{})
	db.AutoMigrate(&model.Family{})
	db.AutoMigrate(&model.Pool{})
	db.AutoMigrate(&model.Meet{})
	db.AutoMigrate(&model.Event{})
	db.AutoMigrate(&model.Entry{})
}

func LoadData(db *gorm.DB) {
	org1, org2 := createOrgs(db)
	createSeasons(org1, db)
	createPersons(org1, org2, db)
}

func createOrgs(db *gorm.DB) (model.Organization, model.Organization) {
	org1 := model.Organization{Label: "Foo Organization"}
	org2 := model.Organization{Label: "Bar Organization"}
	db.Create(&org1)
	db.Create(&org2)

	return org1, org2
}

func createSeasons(org1 model.Organization, db *gorm.DB) (model.Season, model.Season) {
	season1 := model.Season{
		Label: "2020 Season", StartDate: time.Date(2020, 07, 01, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2020, 8, 31, 0, 0, 0, 0, time.UTC),
		Organization: org1,
	}
	season2 := model.Season{
		Label: "2021 Season", StartDate: time.Date(2021, 07, 01, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2021, 8, 31, 0, 0, 0, 0, time.UTC),
		Organization: org1,
	}
	db.Create(&season1)
	db.Create(&season2)

	return season1, season2
}
