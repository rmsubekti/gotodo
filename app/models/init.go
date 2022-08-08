package models

import (
	"gotodo/app/database"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.Connect()
}
