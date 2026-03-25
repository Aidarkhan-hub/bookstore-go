package handlers
import (
    "bookstore/models"
    "encoding/json"
    "net/http"
    "strconv"
)
var Books = []models.Book{}
var NextBookID = 1
func GetBooks(w http.ResponseWriter, r *http.Request) {
    catID := r.URL.Query().Get("category")
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
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
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(filtered[start:end])
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var b models.Book
    if err := json.NewDecoder(r.Body).Decode(&b); err != nil { return }
    b.ID = NextBookID
    NextBookID++
    Books = append(Books, b)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(b)
}
func GetBookByID(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.PathValue("id"))
    for _, b := range Books {
        if b.ID == id {
            json.NewEncoder(w).Encode(b)
            return
        }
    }
    http.Error(w, "Not Found", 404)
}
