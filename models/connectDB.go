package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
)

// DB Pointer
var DB *gorm.DB

// ConnectDB : Tries to connect to the db, exits the program if connection fails
func ConnectDB() {

	// get credentials from .env file
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// open a connection to the database
	db, err := gorm.Open(dbDriver, dbURL)
	if err != nil {
		fmt.Println(fmt.Sprintf("Cannot connect to the databse %+v, error: %+v", dbName, err))
		log.Fatal("Exiting the program")
	}
	fmt.Println("We are connected to the database: ", dbName)

	db.AutoMigrate(&Book{})
	DB = db
}
