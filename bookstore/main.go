package main
import (
    "bookstore/handlers"
    "fmt"
    "net/http"
)
func main() {
    mux := http.NewServeMux()

    // Books
    mux.HandleFunc("GET /books", handlers.GetBooks)
    mux.HandleFunc("GET /books/{id}", handlers.GetBookByID)
    mux.HandleFunc("POST /books", handlers.CreateBook)

    // Authors
    mux.HandleFunc("GET /authors", handlers.GetAuthors)
    mux.HandleFunc("POST /authors", handlers.CreateAuthor)

    // Categories
    mux.HandleFunc("GET /categories", handlers.GetCategories)
    mux.HandleFunc("POST /categories", handlers.CreateCategory)

    fmt.Println("Server starts at :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
