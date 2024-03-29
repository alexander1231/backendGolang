package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn = "host=localhost user=fazt password=mysecretpassword dbname=gorm port=5432"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
