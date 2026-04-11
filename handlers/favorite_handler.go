package handlers

import (
	"bookstore/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var favoriteBooks = []models.FavoriteBook{}

func GetFavorites(c *gin.Context) {
	userID := c.GetInt("user_id")

	firstPg := c.DefaultQuery("L", "1")
	lastPg := c.DefaultQuery("R", "5")

	L, _ := strconv.Atoi(firstPg)
	R, _ := strconv.Atoi(lastPg)

	var userFavorites []models.Book
	for _, fav := range favoriteBooks {
		if fav.UserID == userID {
			for _, b := range books {
				if b.ID == fav.BookID {
					userFavorites = append(userFavorites, b)
				}
			}
		}
	}

	startIndex := (L - 1) * R
	endIndex := startIndex + R

	if startIndex >= len(userFavorites) {
		c.JSON(http.StatusOK, []models.Book{})
		return
	}

	if endIndex > len(userFavorites) {
		endIndex = len(userFavorites)
	}

	c.JSON(http.StatusOK, userFavorites[startIndex:endIndex])
}

func AddFavorite(c *gin.Context) {
	userID := c.GetInt("user_id")
	bookIDStr := c.Param("bookId")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	bookExists := false
	for _, b := range books {
		if b.ID == bookID {
			bookExists = true
			break
		}
	}
	if !bookExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	for _, fav := range favoriteBooks {
		if fav.UserID == userID && fav.BookID == bookID {
			c.JSON(http.StatusConflict, gin.H{"error": "Book already in favorites"})
			return
		}
	}

	newFav := models.FavoriteBook{
		UserID:    userID,
		BookID:    bookID,
		CreatedAt: time.Now(),
	}
	favoriteBooks = append(favoriteBooks, newFav)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Book %d added to favorites", bookID)})
}

func RemoveFavorite(c *gin.Context) {
	userID := c.GetInt("user_id")
	bookIDStr := c.Param("bookId")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for i, fav := range favoriteBooks {
		if fav.UserID == userID && fav.BookID == bookID {
			favoriteBooks = append(favoriteBooks[:i], favoriteBooks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Book %d removed from favorites", bookID)})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found in favorites"})
}
