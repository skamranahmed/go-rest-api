package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skamranahmed/go-rest-api/models"
)

// HandleGetBooks : Get all books (GET /books)
func HandleGetBooks(c *gin.Context) {
	book := models.Book{}

	books, err := book.GetBooks()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, books)
	return
}

// HandleCreateBook : Create a book (POST /books)
func HandleCreateBook(c *gin.Context) {
	var input models.CreateBookInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	book := &models.Book{}
	book, err = book.CreateBook()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
	return
}
