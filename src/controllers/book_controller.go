package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hokkyar/rest-bookapp/src/config"
	"github.com/hokkyar/rest-bookapp/src/models"
)

func GetBooks(c *gin.Context) {
	var book []models.Book

	if err := config.GetDB().Preload("User").Find(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := config.GetDB().Preload("User").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func PostBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.GetDB().First(&user, book.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := config.GetDB().Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success create Book",
		"book":    book,
	})
}

func PutBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := config.GetDB().First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		Year        int    `json:"year"`
		Description string `json:"description"`
		UserID      uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Name = input.Name
	book.Publisher = input.Publisher
	book.Year = input.Year
	book.Description = input.Description
	book.UserID = input.UserID

	var user models.User
	if err := config.GetDB().First(&user, book.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := config.GetDB().Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}

	response := struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		Year        int    `json:"year"`
		Description string `json:"description"`
	}{
		ID:          book.ID,
		Name:        book.Name,
		Publisher:   book.Publisher,
		Year:        book.Year,
		Description: book.Description,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update book",
		"book":    response,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := config.GetDB().First(&models.Book{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := config.GetDB().Delete(&models.Book{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting book"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
