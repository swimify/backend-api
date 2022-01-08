package fixtures

import (
	"github.com/swimlibre/backend-api/model"
	"gorm.io/gorm"
)

func loadAddresses(db *gorm.DB) {
	a1 := model.Address{}
	db.Create(&a1)
}
