package controllers

import (
	"github.com/nikola43/ethfaucet/models"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&models.User{})

	// CREATE
	GormDB.AutoMigrate(&models.User{})
}
