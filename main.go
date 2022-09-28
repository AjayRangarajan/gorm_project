package main

import (
	"fmt"
	"gorm_project/database"
	"gorm_project/routes"
	"log"
	"net/http"

	"github.com/subosito/gotenv"
)

func init() {
	env_err := gotenv.Load()
	if env_err != nil {
		log.Panic(env_err)
	}
}

func main() {
	fmt.Println("0:1")
	var db = database.Db()
	fmt.Println("0:2")
	if db.Error != nil {
		fmt.Println("0:3")
		fmt.Println(db.Error)
	} else {
		fmt.Println("Successfully created tables.")
	}

	fmt.Println("0:4")
	var appRouter = routes.CreateRouter()
	fmt.Println("0:5")
	log.Println("Listening on Port 8000")
	fmt.Println("0:6")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
	fmt.Println("0:7")
}
