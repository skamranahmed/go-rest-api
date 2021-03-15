package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

	book := &models.Book{Title: input.Title, Author: input.Author}
	book, err = book.CreateBook()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
	return
}

// HandleGetBook : Get a book (GET /books/:id)
func HandleGetBook(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	book, err := models.GetBook(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
	return
}

// HandleUpdateBook : Update a book (PUT /books/:id)
func HandleUpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	book, err := models.GetBook(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var input models.UpdateBookInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	book = &models.Book{ID: book.ID, Title: input.Title, Author: input.Author}
	book, err = book.UpdateBook()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, book)
	return
}
