// Router for the app
package routes

import (
	"gorm_project/services"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/companies", services.GetAllCompanies).Methods("GET")
	router.HandleFunc("/company", services.CreateCompany).Methods("POST")

	return router
}
