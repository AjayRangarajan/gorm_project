package main

import (
	"fmt"
	"gorm_project/database"
)

func main() {
	var db = database.Db()
	if db.Error != nil {
		fmt.Println(db.Error)
	} else {
		fmt.Println("Successfully created tables.")
	}
}
