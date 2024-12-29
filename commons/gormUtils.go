package commons

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	m"messagingapp/models"
)

func InitGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("messagingapp.db"), &gorm.Config{})
	if err != nil {
		panic("cannot connect to gorm: " + err.Error())
	}
	db.AutoMigrate(&m.User{})

	return db
}