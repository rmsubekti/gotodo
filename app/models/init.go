package models

import (
	"gotodo/app/config"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.GetDB()

	db.AutoMigrate(&User{}, &Project{}, &Task{})
}
