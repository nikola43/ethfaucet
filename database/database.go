package controllers

import (
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&models.Client{})

	// CREATE
	GormDB.AutoMigrate(&models.Client{})
}

