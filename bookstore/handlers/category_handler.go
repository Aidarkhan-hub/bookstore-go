package handlers
import ("bookstore/models"; "encoding/json"; "net/http")

var Categories = []models.Category{}
var NextCategoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
    var c models.Category
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    c.ID = NextCategoryID
    NextCategoryID++
    Categories = append(Categories, c)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(c)
}
