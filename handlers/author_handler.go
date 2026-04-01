package handlers

import (
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.Author{
	{ID: 1, Name: "Aidarkhan"},
	{ID: 2, Name: "Akzhan"},
	{ID: 3, Name: "Batyrkhan"},
	{ID: 4, Name: "Bauyrzhan"},
	{ID: 5, Name: "Sholpan"},
}

func GetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, users)

}

func AddAuthor(c *gin.Context) {
	var user models.Author

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	user.ID = len(users) + 1
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}
