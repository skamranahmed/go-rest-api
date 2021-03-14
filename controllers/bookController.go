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
