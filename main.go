package main

import (
	"API/database"
	"API/models"
	"API/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer database.DB.Close()
	database.DB.AutoMigrate(&models.ModelForm{})
	r := routes.SetupRouter()

	//running
	r.Run(":5000")
}
