package handlers

import (
    "net/http"
	"strconv"
    "github.com/gin-gonic/gin"
    "github.com/baselrabia/book-microservice/models"
)

var books []models.Book

// CreateBook adds a new book to the store
func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    books = append(books, book)
    c.JSON(http.StatusCreated, book)
}

// UpdateBook updates a book by ID
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook models.Book
    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find and update the book with the given ID
    for i, book := range books {
        if strconv.Itoa(book.ID) == id {
            books[i] = updatedBook
            c.JSON(http.StatusOK, updatedBook)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// DeleteAllBooks deletes all books
func DeleteAllBooks(c *gin.Context) {
    books = []models.Book{}
    c.JSON(http.StatusOK, gin.H{"message": "All books deleted"})
}

// GetBooks retrieves all books
func GetBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

// GetBookByID retrieves a specific book by ID
func GetBookByID(c *gin.Context) {
    id := c.Param("id")

    for _, book := range books {
        if strconv.Itoa(book.ID) == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
