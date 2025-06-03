package main

import (
    "encoding/json"
    "net/http"
)

type Post struct {
    Name    string `json:"name"`
    Message string `json:"message"`
}

var posts = []Post{
    {"Alice", "Hello!"},
    {"Bob", "Hi!"},
}

func main() {
    http.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(posts)
    })

    http.HandleFunc("/api/post", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            name := r.FormValue("name")
            msg := r.FormValue("message")
            posts = append(posts, Post{name, msg})
            w.Write([]byte("OK"))
        }
    })

    http.ListenAndServe(":8080", nil)
}

