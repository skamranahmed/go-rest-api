package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skamranahmed/go-rest-api/controllers"
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
	r.GET("/books", controllers.HandleGetBooks)
	r.POST("/books", controllers.HandleCreateBook)
	r.GET("/books/:id", controllers.HandleGetBook)
	r.PUT("/books/:id", controllers.HandleUpdateBook)
	r.DELETE("/books/:id", controllers.HandleDeleteBook)

	r.Run(":8000")
}
