package database

import (
	"gorm_project/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.Company{})
	db.AutoMigrate(&models.CompanyPhoneNumber{})
	db.AutoMigrate(&models.Workflow{})
	db.AutoMigrate(&models.Stage{})
	db.AutoMigrate(&models.Campaign{})
	db.AutoMigrate(&models.CampaignLog{})
	db.AutoMigrate(&models.Position{})
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.PhoneNumber{})
	db.AutoMigrate(&models.Email{})

	return db
}
