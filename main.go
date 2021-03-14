package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skamranahmed/go-rest-api/models"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error getting env variables: ", err)
		log.Fatalf("Exiting the program")
	}

	// initializing default gin router
	r := gin.Default()

	// connect db
	models.ConnectDB()

	r.Run(":8000")
}
