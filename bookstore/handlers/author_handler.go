package handlers
import ("bookstore/models"; "encoding/json"; "net/http")
var Authors = []models.Author{}
var NextAuthorID = 1
func GetAuthors(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Authors)
}
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
    var a models.Author
    json.NewDecoder(r.Body).Decode(&a)
    a.ID = NextAuthorID
    NextAuthorID++
    Authors = append(Authors, a)
    json.NewEncoder(w).Encode(a)
}
