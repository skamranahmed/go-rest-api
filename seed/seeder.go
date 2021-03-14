package seed

import (
	"log"

	"github.com/skamranahmed/go-rest-api/models"
)

var books = []models.Book{
	models.Book{
		ID:     1,
		Title:  "Book 1",
		Author: "Author 1",
	},
	models.Book{
		ID:     2,
		Title:  "Book 2",
		Author: "Author 2",
	},
}

// PopulateDB : populates db with dummy seed data
func PopulateDB() {
	err := models.DB.DropTableIfExists(&models.Book{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %+v", err)
	}

	err = models.DB.AutoMigrate(&models.Book{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %+v", err)
	}

	for i := range books {
		err := models.DB.Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}
	}
}
