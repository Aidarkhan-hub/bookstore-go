package main

import (
    "bookstore/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/books", handlers.GetBooks)
    r.GET("/books/:id", handlers.GetBookByID)
    r.POST("/books", handlers.CreateBook)

    // Для авторов и категорий (добавь аналогично в handlers)
    // r.GET("/authors", handlers.GetAuthors)
    // r.POST("/authors", handlers.CreateAuthor)

    r.Run(":8080")
}
