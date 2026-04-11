package main

import (
	"bookstore/handlers"
	"bookstore/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", handlers.GetAuthors)
	r.POST("/users", handlers.AddAuthor)

	r.GET("/books", handlers.GetBook)
	r.POST("/books", handlers.AddBook)
	r.PUT("/books/:bookId", handlers.UpdateBook)
	r.GET("/books/:bookId", handlers.GetBookByID)
	r.DELETE("/books/:bookId", handlers.DeleteBook)

	r.GET("/categories", handlers.GetCategory)
	r.POST("/categories", handlers.AddCategory)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/books/favorites", handlers.GetFavorites)
		auth.PUT("/books/:bookId/favorites", handlers.AddFavorite)
		auth.DELETE("/books/:bookId/favorites", handlers.RemoveFavorite)
	}

	r.Run(":8080")
}
