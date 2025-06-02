package main

import (
    "fmt"
    "html/template"
    "net/http"
    "sync"
)

type Post struct {
    Name    string
    Message string
}

var posts []Post
var mu sync.Mutex

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/submit", submitHandler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}

// コメント見やすくて草
func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, posts)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        message := r.FormValue("message")
        mu.Lock()
        posts = append(posts, Post{Name: name, Message: message})
        mu.Unlock()
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}
