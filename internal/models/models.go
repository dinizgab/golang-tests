package models

type User struct {
    ID        string 
    FirstName string
    Username  string
    Posts     []Post
}

type Post struct {
    ID        string 
    Title     string
    Body      string
}
