package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Id            int `gorm:"primaryKey;autoIncrement:true"`
	Name          string
	URL           string
	EmployeeCount int
	LinkedInURL   string
}

type CompanyPhoneNumber struct {
	gorm.Model
	Id            int     `gorm:"primaryKey;autoIncrement:true"`
	Company       Company `gorm:"foreignKey:Id"`
	PhoneNumber   string
	PhoneType     string
	CountryCode   string
	LastValidated time.Time
}

type Workflow struct {
	gorm.Model
	Id           int `gorm:"primaryKey;autoIncrement:true"`
	Name         string
	UUID         string
	Organization string
	Team         string
}

type Stage struct {
	gorm.Model
	Id              int `gorm:"primaryKey;autoIncrement:true"`
	VerticalOrder   int
	HorizontalOrder int
	Workflow        Workflow `gorm:"foreignKey:Id"`
	Name            string
	Finished        bool
}

type Campaign struct {
	gorm.Model
	Id   int `gorm:"primaryKey;autoIncrement:true"`
	Name string
}

type CampaignLog struct {
	gorm.Model
	Id       int      `gorm:"primaryKey;autoIncrement:true"`
	Campaign Campaign `gorm:"foreignKey:Id"`
	Person   string
	Action   string
}

type Position struct {
	gorm.Model
	Id   int `gorm:"primaryKey;autoIncrement:true"`
	Name string
}

type Person struct {
	gorm.Model
	Id          int `gorm:"primaryKey;autoIncrement:true"`
	FirstName   string
	LastName    string
	Company     Company  `gorm:"foreignKey:Id"`
	Position    Position `gorm:"foreignKey:Id"`
	LinkedInURL string
}

type PhoneNumber struct {
	gorm.Model
	Id            int    `gorm:"primaryKey;autoIncrement:true"`
	Person        Person `gorm:"foreignKey:Id"`
	PhoneNumber   string
	PhoneType     string
	LastValidated time.Time
	CountryCode   string
}

type Email struct {
	gorm.Model
	Id     int    `gorm:"primaryKey;autoIncrement:true"`
	Person Person `gorm:"foreignKey:Id"`
	Value  string
	Type   string
}
