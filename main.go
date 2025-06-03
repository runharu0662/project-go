package main

import (
    "fmt"
    "net/http"
)

func handlePost(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    name := r.FormValue("name")
    message := r.FormValue("message")
    fmt.Printf("投稿: %s - %s\n", name, message)
    w.Write([]byte("投稿を受け付けました"))
}

func main() {
    http.HandleFunc("/api/post", handlePost)
    fmt.Println("Go API running on :8080")
    http.ListenAndServe(":8080", nil)
}

