package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skamranahmed/go-rest-api/handler"
	"github.com/skamranahmed/go-rest-api/models"
)

func main() {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error getting env variables: ", err)
		log.Fatalf("Exiting the program")
	}

	// initializing default gin router
	r := gin.Default()

	// connect to db
	models.ConnectDB()

	// // seed dummy data in db
	// seed.PopulateDB()

	// routes
	r.GET("/books", handler.HandleGetBooks)
	r.POST("/books", handler.HandleCreateBook)
	r.GET("/books/:id", handler.HandleGetBook)
	r.PUT("/books/:id", handler.HandleUpdateBook)
	r.DELETE("/books/:id", handler.HandleDeleteBook)

	r.Run(":8000")
}
