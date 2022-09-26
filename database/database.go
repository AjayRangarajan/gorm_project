package database

import (
	"fmt"
	"gorm_project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=gorm_project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
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
