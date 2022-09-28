package services

import (
	"encoding/json"
	"fmt"
	"gorm_project/models"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

type Response struct {
	Data    []models.Company `json:"data"`
	Message string           `json:"message"`
}

func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("2")
	var companies = models.GetCompanies()
	fmt.Println("3")
	var res Response
	err := db.Find(&companies).Error
	fmt.Println("4")
	if err == nil {
		log.Println(companies)
		res.Data = companies
		fmt.Println("5")
		res.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&res)
		fmt.Println("6")
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res Response
	var company = models.GetCompany()
	_ = json.NewDecoder(r.Body).Decode(&company)
	log.Println(company)
	err := db.Create(&company).Error
	if err != nil {
		http.Error(w, "Error Creating Record", 400)
		return
	}
	res.Message = "CREATED"
	json.NewEncoder(w).Encode(res)
}
