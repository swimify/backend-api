package fixtures

import (
	"github.com/swimlibre/backend-api/model"
	"gorm.io/gorm"
)

func createPersons(org1 model.Organization, org2 model.Organization, db *gorm.DB) []model.Person {
	// COACHES
	p1 := model.Person{FirstName: "Org 1, Coach #1", Organization: org1, IsCoach: true}
	db.Create(&p1)

	p2 := model.Person{FirstName: "Org 1, Coach #2", Organization: org1, IsCoach: true}
	db.Create(&p2)

	// FAMILY #1
	p3 := model.Person{FirstName: "Org 1, Swimmer #1 Family #1", Organization: org1, IsSwimmer: true}
	db.Create(&p3)

	p4 := model.Person{FirstName: "Org 1, Swimmer #2 Family #1", Organization: org1, IsSwimmer: true}
	db.Create(&p4)

	p5 := model.Person{FirstName: "Org 1, Guardian #1, Family #1", Organization: org1, IsGuardian: true}
	db.Create(&p5)

	f1 := model.Family{Persons: []model.Person{p3, p4, p5}, Organization: org1}
	db.Create(&f1)

	// FAMILY #2
	p6 := model.Person{FirstName: "Org 1, Swimmer #1 Family #2", Organization: org1, IsSwimmer: true}
	db.Create(&p6)

	p7 := model.Person{FirstName: "Org 1, Swimmer #2 Family #2", Organization: org1, IsSwimmer: true}
	db.Create(&p7)

	p8 := model.Person{FirstName: "Org 1, Guardian #1, Family #2", Organization: org1, IsGuardian: true}
	db.Create(&p8)

	p9 := model.Person{FirstName: "Org 1, Guardian #1, Family #2", Organization: org1, IsGuardian: true}
	db.Create(&p9)

	f2 := model.Family{Persons: []model.Person{p6, p7, p8, p9}, Organization: org1}
	db.Create(&f2)

	// CREATE People in Org #2
	p10 := model.Person{FirstName: "Org 2, Coach #1", Organization: org2, IsCoach: true}
	db.Create(&p10)
	p11 := model.Person{FirstName: "Org 2, Swimmer #1", Organization: org2, IsSwimmer: true}
	db.Create(&p11)

	return []model.Person{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11}
}
