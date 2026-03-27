package handlers

import (
    "bookstore/models"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

var Books = []models.Book{}
var NextBookID = 1

func GetBooks(c *gin.Context) {
    catID := c.Query("category")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit := 5
    
    var filtered []models.Book
    for _, b := range Books {
        if catID == "" || strconv.Itoa(b.CategoryID) == catID {
            filtered = append(filtered, b)
        }
    }

    start := (page - 1) * limit
    if start < 0 || start >= len(filtered) { start = 0 }
    end := start + limit
    if end > len(filtered) { end = len(filtered) }

    c.JSON(http.StatusOK, filtered[start:end])
}

func CreateBook(c *gin.Context) {
    var b models.Book
    if err := c.ShouldBindJSON(&b); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    b.ID = NextBookID
    NextBookID++
    Books = append(Books, b)
    c.JSON(http.StatusCreated, b)
}

func GetBookByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, b := range Books {
        if b.ID == id {
            c.JSON(http.StatusOK, b)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
