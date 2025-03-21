package database

import (
	"fmt"
	"log"
	
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := `host=localhost user=postgres password=1
	dbname=users port=5432 sslmode=disable TimeZone=Asia/Tehran`

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error in connect database, error = %v \n", err)
	}

	DB = db
	fmt.Println("connect database successfully")
}
