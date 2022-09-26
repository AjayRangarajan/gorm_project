package main

import (
	"fmt"
	"gorm_project/database"
	"log"

	"github.com/subosito/gotenv"
)

func init() {
	env_err := gotenv.Load()
	if env_err != nil {
		log.Panic(env_err)
	}
}

func main() {
	var db = database.Db()
	if db.Error != nil {
		fmt.Println(db.Error)
	} else {
		fmt.Println("Successfully created tables.")
	}
}
