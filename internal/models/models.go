package models

type User struct {
    ID        string `json:"id"`
    FirstName string `json:"name"`
    Username  string `json:"username"`
    Posts     []Post `json:"posts"`
}

type Post struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Body      string `json:"body"`
}
