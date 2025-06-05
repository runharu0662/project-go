package main

import (
	"encoding/json"
	"fmt"
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
	// cout "Server is running on http://localhost:8080"
	fmt.Println("Server is running on http://localhost:8080")
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
